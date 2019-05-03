package topdesk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"sync"
)

func (c *Client) list(ctx context.Context, resource string) (*ListIterator, error) {
	return &ListIterator{
		client:   c,
		ctx:      ctx,
		more:     true,
		resource: resource,
		data:     make([]json.RawMessage, 0),
	}, nil
}

type ListIterator struct {
	client   *Client
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

func (l *ListIterator) More() bool {
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
