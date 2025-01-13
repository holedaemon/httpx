package httpx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func newRequest(opts ...RequestOption) (*Request, error) {
	r := new(Request)
	for _, o := range opts {
		o(r)
	}

	if r.method == "" {
		return nil, errors.New("httpx: method not set")
	}

	if r.url == "" {
		return nil, errors.New("httpx: url not set")
	}

	if !strings.HasPrefix(r.url, "/") {
		r.url = "/" + r.url
	}

	return r, nil
}

// Do performs an HTTP request using the given options and returns the
// resulting *[http.Response]. It is up to consumers to perform error handling
// on the returned status, as well as close the response body when finished.
func (c *Client) Do(ctx context.Context, opts ...RequestOption) (*http.Response, error) {
	r, err := newRequest(opts...)
	if err != nil {
		return nil, fmt.Errorf("preparing request: %w", err)
	}

	u := fmt.Sprintf("%s%s", c.rootURL, r.url)
	req, err := http.NewRequestWithContext(ctx, r.method, u, r.body)
	if err != nil {
		return nil, fmt.Errorf("preparing request: %w", err)
	}

	if r.header != nil {
		req.Header = r.header.Clone()
	}

	req.Header.Set("User-Agent", c.userAgent)

	if r.query != nil {
		req.URL.RawQuery = r.query.Encode()
	}

	res, err := c.c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("performing request: %w", err)
	}

	return res, nil
}
