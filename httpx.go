// Package httpx provides useful extensions to the net/http standard library.
package httpx

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

var (
	// DefaultClient is an *[http.Client] with sane defaults.
	DefaultClient = &http.Client{
		Timeout: time.Second * 5,
	}

	// ErrStatus is returned when a *[Client] receives a status code outside of
	// the 200 range.
	ErrStatus = errors.New("httpx: server returned an unexpected response code")

	userAgent = fmt.Sprintf("%s/v%s", Slug, Version)
)

// Client is a thin wrapper over *[http.Client], providing a convenient API and
// defaults for making requests with.
type Client struct {
	userAgent string
	rootURL   string

	c *http.Client
}

// New creates a new *[Client] from the given options. If an *[http.Client]
// is not provided, [DefaultClient] is used.
func New(opts ...Option) (*Client, error) {
	c := new(Client)
	for _, o := range opts {
		o(c)
	}

	if c.userAgent == "" {
		c.userAgent = userAgent
	}

	if c.rootURL == "" {
		return nil, errors.New("httpx: root URL cannot be blank")
	}

	c.rootURL = strings.TrimSuffix(c.rootURL, "/")

	if c.c == nil {
		c.c = DefaultClient
	}

	return c, nil
}
