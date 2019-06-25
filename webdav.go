package topdesk

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// Webdav client.
type WebdavClient struct {
	endpoint      *url.URL
	authorization string
}

// New Webdav client.
func NewWebdavClient(ctx context.Context, endpoint string, authorization string) (*WebdavClient, error) {
	uri, err := url.Parse(endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "parse endpoint url")
	}

	if _, err := base64.StdEncoding.DecodeString(authorization); err != nil {
		authorization = base64.StdEncoding.EncodeToString([]byte(authorization))
	}

	wc := &WebdavClient{
		endpoint:      uri,
		authorization: fmt.Sprintf("Basic %s", authorization),
	}

	// TODO(shane): Test connection.
	return wc, nil
}

// Put a file on the server.
func (wc *WebdavClient) Put(context context.Context, filepath string, file io.Reader) error {
	parts := strings.Split(filepath, "?")

	uri := *wc.endpoint
	uri.Path = path.Join(uri.Path, parts[0])
	if len(parts) > 1 {
		uri.RawQuery = parts[1]
	}

	req, _ := http.NewRequest(http.MethodPut, uri.String(), file)
	req.Header.Add("Authorization", wc.authorization)
	req.Header.Add("Content-Type", "binary/octet-stream")

	// TODO(shane): Dialer timeout.
	client := &http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		errors.Wrapf(err, "%s put %s", http.StatusText(res.StatusCode), uri.String())
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated, http.StatusNoContent:
		return nil
	default:
		return fmt.Errorf("%s put %s", http.StatusText(res.StatusCode), uri.String())
	}
}
