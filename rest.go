package topdesk

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
)

// Rest client.
type RestClient struct {
	endpoint      *url.URL
	authorization string
}

// New REST client.
func NewRestClient(ctx context.Context, endpoint string, authorization string) (*RestClient, error) {
	uri, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "parse endpoint url")
	}

	if _, err := base64.StdEncoding.DecodeString(authorization); err != nil {
		authorization = base64.StdEncoding.EncodeToString([]byte(authorization))
	}

	rc := &RestClient{
		endpoint:      uri,
		authorization: fmt.Sprintf("Basic %s", authorization),
	}

	// TODO(shane): Test connection.
	return rc, nil
}

func (rc *RestClient) do(context context.Context, method string, uri *url.URL, request interface{}, response interface{}) (int, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return http.StatusBadRequest, errors.Wrapf(err, "%s %s encoding request body", method, uri.String())
	}

	req, _ := http.NewRequest(method, uri.String(), bytes.NewReader(body))
	req.Header.Add("Authorization", rc.authorization)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	// debugging, _ := httputil.DumpRequest(req, true)
	// fmt.Printf("%s\n", debugging)

	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return http.StatusBadRequest, errors.Wrapf(err, "%s %s", method, uri.String())
	}
	defer res.Body.Close()

	// debugging, _ = httputil.DumpResponse(res, true)
	// fmt.Printf("%s\n", debugging)

	switch res.StatusCode {
	case http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError:
		// Return a
		messages := ErrorMessages{}
		if err := json.NewDecoder(res.Body).Decode(&messages); err != nil {
			return http.StatusBadRequest, errors.Wrapf(err, "%s %s decoding response body", method, uri.String())
		}
		return res.StatusCode, messages

	case http.StatusOK, http.StatusCreated, http.StatusPartialContent:
		if err := json.NewDecoder(res.Body).Decode(response); err != nil {
			return http.StatusBadRequest, errors.Wrapf(err, "%s %s decoding response body", method, uri.String())
		}
	}

	return res.StatusCode, nil
}

func (rc *RestClient) get(ctx context.Context, endpoint *url.URL, response interface{}) error {
	status, err := rc.do(ctx, http.MethodGet, endpoint, nil, response)
	switch {
	case err != nil:
		return err
	case status == http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s save %s", http.StatusText(status), endpoint.String())
	}
}

func (rc *RestClient) create(ctx context.Context, endpoint *url.URL, request interface{}, response interface{}) error {
	status, err := rc.do(ctx, http.MethodPost, endpoint, request, response)
	switch {
	case err != nil:
		return err
	case status == http.StatusOK || status == http.StatusCreated:
		return nil
	default:
		return fmt.Errorf("%s save %s", http.StatusText(status), endpoint.String())
	}
}

func (rc *RestClient) update(ctx context.Context, endpoint *url.URL, request interface{}, response interface{}) error {
	status, err := rc.do(ctx, http.MethodPut, endpoint, request, response)
	switch {
	case err != nil:
		return err
	case status == http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s save %s", http.StatusText(status), endpoint.String())
	}
}

func (rc *RestClient) delete(ctx context.Context, endpoint *url.URL) error {
	status, err := rc.do(ctx, http.MethodDelete, endpoint, nil, nil)
	switch {
	case err != nil:
		return err
	case status == http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s save %s", http.StatusText(status), endpoint.String())
	}
}

func (rc *RestClient) list(ctx context.Context, endpoint *url.URL) (*ListIterator, error) {
	return &ListIterator{
		start:    0,
		pageSize: 100, // Magic number, but it's Topdesk max.
		client:   rc,
		ctx:      ctx,
		more:     true,
		endpoint: endpoint,
		data:     make([]json.RawMessage, 0),
	}, nil
}

type ListIterator struct {
	client   *RestClient
	start    uint64
	pageSize uint64
	more     bool
	endpoint *url.URL
	ctx      context.Context
	mu       sync.Mutex
	data     []json.RawMessage
}

func (l *ListIterator) decode(response interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if len(l.data) == 0 {
		return errors.New("no data")
	}

	err := json.Unmarshal(l.data[0], response)
	if err == nil {
		l.data = l.data[1:]
	}
	return err
}

func (l *ListIterator) Next() bool {
	if len(l.data) == 0 && l.more {
		uri := *l.endpoint

		query := uri.Query()
		query.Set("page_size", fmt.Sprintf("%d", l.pageSize))
		query.Set("start", fmt.Sprintf("%d", l.start))
		uri.RawQuery = query.Encode()

		status, err := l.client.do(l.ctx, http.MethodGet, &uri, nil, &l.data)
		if err != nil {
			return false
		}
		l.start = l.start + l.pageSize
		l.more = (status == http.StatusPartialContent)
	}

	return len(l.data) > 0
}

// ErrorMessages REST API response.
type ErrorMessages []struct {
	Message string `json:"message"`
}

func (e ErrorMessages) Error() string {
	errs := []string{}
	for _, em := range e {
		errs = append(errs, em.Message)
	}
	return strings.Join(errs, " ")
}

// Ref is a resource reference.
type Ref struct {
	ID string `"id"`
}

// ResourceRef creates a reference from any resource that implements the interface.
//
// A resource returned by Get* may not be compatible with an Update*. Often the request object only accepts
// a Ref struct not the ID string or the complete object.
type ResourceRef interface {
	Ref() Ref
}
