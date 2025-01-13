package httpx

import (
	"io"
	"net/http"
	"net/url"
)

// Request represents an HTTP request.
type Request struct {
	method string
	url    string
	body   io.Reader
	header http.Header
	query  url.Values
}

// RequestOption is used to configure an HTTP request.
type RequestOption func(*Request)

// URL configures the URL of a request. A leading / character will be added
// if not already present.
func URL(u string) RequestOption {
	return func(r *Request) {
		r.url = u
	}
}

// Body configures a request to use the given [io.Reader] as a body.
func Body(rdr io.Reader) RequestOption {
	return func(r *Request) {
		r.body = rdr
	}
}

// Header adds the given key and value to a request's headers.
func Header(key, val string) RequestOption {
	return func(r *Request) {
		if r.header == nil {
			r.header = make(http.Header)
		}

		r.header.Set(key, val)
	}
}

// Query adds the given key and value to a request's query.
func Query(key, val string) RequestOption {
	return func(r *Request) {
		if r.query == nil {
			r.query = make(url.Values)
		}

		r.query.Set(key, val)
	}
}

// Get sets a request's method to GET.
func Get() RequestOption {
	return func(r *Request) {
		r.method = http.MethodGet
	}
}

// Head sets a request's method to HEAD.
func Head() RequestOption {
	return func(r *Request) {
		r.method = http.MethodHead
	}
}

// Post sets a request's method to POST.
func Post() RequestOption {
	return func(r *Request) {
		r.method = http.MethodPost
	}
}

// Put sets a request's method to PUT.
func Put() RequestOption {
	return func(r *Request) {
		r.method = http.MethodPut
	}
}

// Delete sets a request's method to DELETE.
func Delete() RequestOption {
	return func(r *Request) {
		r.method = http.MethodDelete
	}
}

// Connect sets a request's method to CONNECT.
func Connect() RequestOption {
	return func(r *Request) {
		r.method = http.MethodConnect
	}
}

// Options sets a request's method to OPTIONS.
func Options() RequestOption {
	return func(r *Request) {
		r.method = http.MethodOptions
	}
}

// Trace sets a request's method to TRACE.
func Trace() RequestOption {
	return func(r *Request) {
		r.method = http.MethodTrace
	}
}

// Patch sets a request's method to PATCH.
func Patch() RequestOption {
	return func(r *Request) {
		r.method = http.MethodPatch
	}
}
