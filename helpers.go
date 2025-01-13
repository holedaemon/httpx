package httpx

import "net/http"

// OK determines if the given status is within the success range.
func OK(status int) bool {
	return status >= http.StatusOK && status < http.StatusMultipleChoices
}
