package httpx

import "net/http"

// Option is used to configure a *[Client].
type Option func(*Client)

// HTTPClient sets the underlying *[http.Client] of a *[Client].
func HTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.c = client
	}
}

// UserAgent sets the default User-Agent of a *[Client].
func UserAgent(ua string) Option {
	return func(c *Client) {
		c.userAgent = ua
	}
}

// RootURL sets the root URL of a *[Client].
func RootURL(ru string) Option {
	return func(c *Client) {
		c.rootURL = ru
	}
}
