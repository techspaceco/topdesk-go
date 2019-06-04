package topdesk

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
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

func (rc *RestClient) do(context context.Context, method string, uri string, request interface{}, response interface{}) (int, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return http.StatusBadRequest, err
	}

	body, err := json.Marshal(request)
	if err != nil {
		return http.StatusBadRequest, errors.Wrapf(err, "%s %s encoding request body", method, u.String())
	}

	req, _ := http.NewRequest(method, u.String(), bytes.NewReader(body))
	req.Header.Add("Authorization", rc.authorization)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json;charset=utf-8")

	debugging, _ := httputil.DumpRequest(req, true)
	fmt.Printf("%s\n", debugging)

	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return http.StatusBadRequest, errors.Wrapf(err, "%s %s", method, u.String())
	}
	defer res.Body.Close()

	debugging, _ = httputil.DumpResponse(res, true)
	fmt.Printf("%s\n", debugging)

	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return http.StatusBadRequest, errors.Wrapf(err, "%s %s decoding response body", method, u.String())
	}

	return res.StatusCode, nil
}

func (rc *RestClient) get(ctx context.Context, resource string, id string, response interface{}) error {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, resource)
	if len(id) > 0 {
		uri.Path = path.Join(uri.Path, "id", id)
	}

	status, err := rc.do(ctx, http.MethodGet, uri.String(), nil, response)
	if err != nil {
		return err
	}

	switch status {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s delete %s", http.StatusText(status), uri.String())
	}
}

func (rc *RestClient) save(ctx context.Context, resource string, id string, request interface{}, response interface{}) error {
	method := http.MethodPost
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, resource)
	if len(id) > 0 {
		uri.Path = path.Join(uri.Path, "id", id)
	}

	status, err := rc.do(ctx, method, uri.String(), request, response)
	if err != nil {
		return err
	}

	switch status {
	case http.StatusCreated, http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s save %s", http.StatusText(status), uri.String())
	}
}

func (rc *RestClient) delete(ctx context.Context, resource string, id string) error {
	uri := *rc.endpoint
	uri.Path = path.Join(uri.Path, resource)
	if len(id) > 0 {
		uri.Path = path.Join(uri.Path, "id", id)
	}

	status, err := rc.do(ctx, http.MethodDelete, uri.String(), nil, nil)
	if err != nil {
		return err
	}

	switch status {
	case http.StatusOK:
		return nil
	default:
		return fmt.Errorf("%s delete %s", http.StatusText(status), uri.String())
	}
}

func (rc *RestClient) list(ctx context.Context, resource string) (*ListIterator, error) {
	return &ListIterator{
		client:   rc,
		ctx:      ctx,
		more:     true,
		resource: resource,
		data:     make([]json.RawMessage, 0),
	}, nil
}

type ListIterator struct {
	client   *RestClient
	start    uint64
	pageSize uint64
	more     bool
	resource string
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
		uri := *l.client.endpoint
		uri.Path = path.Join(uri.Path, l.resource)

		start := uint64(0)
		if l.start > 0 {
			start = l.start + l.pageSize
		}

		query := url.Values{}
		query.Set("page_size", "100")
		query.Set("start", fmt.Sprintf("%d", start))
		uri.RawQuery = query.Encode()

		status, err := l.client.do(l.ctx, http.MethodGet, uri.String(), nil, &l.data)
		if err != nil {
			return false
		}
		l.more = (status == http.StatusPartialContent)
	}

	return len(l.data) > 0
}
