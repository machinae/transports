package transports

import (
	"context"
	"net/http"
)

// HeaderTransport sets the given Headers on each outbound request
// It does not overwrite headers with the same key already set on the request
type HeaderTransport struct {
	http.RoundTripper

	Headers map[string]string
}

func (tr *HeaderTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if tr.RoundTripper == nil {
		tr.RoundTripper = http.DefaultTransport
	}
	// Make a shallow copy of the request before modifying
	req = req.WithContext(context.Background())
	// Since req.Header is a reference, make a new one and copy the values
	headers := make(http.Header)
	for k, v := range tr.Headers {
		headers.Set(k, v)
	}
	// Apply original request headers, overwriting any set by the transport
	for k := range req.Header {
		headers[k] = append([]string{}, req.Header[k]...)
	}
	req.Header = headers
	return tr.RoundTripper.RoundTrip(req)
}
