package topdesk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// Client ...
type Client struct {
	endpoint      *url.URL
	authorization string
}

// New ...
func New(ctx context.Context, endpoint string, authorization string) (*Client, error) {
	uri, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "parse endpoint url")
	}

	c := &Client{
		endpoint:      uri,
		authorization: fmt.Sprintf("Basic %s", authorization),
	}

	// TODO: Test connection.
	return c, nil
}

func (c *Client) do(context context.Context, method string, uri string, request interface{}, response interface{}) (int, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return http.StatusBadRequest, err
	}

	body, err := json.Marshal(request)
	if err != nil {
		return http.StatusBadRequest, errors.Wrapf(err, "%s %s encoding request body", method, u.String())
	}

	req, _ := http.NewRequest(method, u.String(), bytes.NewReader(body))
	req.Header.Add("Authorization", c.authorization)
	req.Header.Add("Accept", "application/json")

	// debugging, _ := httputil.DumpRequest(req, true)
	// fmt.Printf("%s\n", debugging)

	client := &http.Client{Timeout: 5 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return res.StatusCode, errors.Wrapf(err, "%s %s", method, u.String())
	}
	defer res.Body.Close()

	// debugging, _ = httputil.DumpResponse(res, true)
	// fmt.Printf("%s\n", debugging)

	if err := json.NewDecoder(res.Body).Decode(response); err != nil {
		return http.StatusBadRequest, errors.Wrapf(err, "%s %s decoding response body", method, u.String())
	}

	return res.StatusCode, nil
}
