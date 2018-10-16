package transports

import (
	"context"
	"net/http"
	"time"
)

//TimeoutTransport enforces a timeout on all requests
type TimeoutTransport struct {
	http.RoundTripper
	// Timeout on all requests with this transport
	Timeout time.Duration
}

func (tr *TimeoutTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if tr.RoundTripper == nil {
		tr.RoundTripper = http.DefaultTransport
	}
	if tr.Timeout == 0 {
		tr.Timeout = DefaultTimeout
	}
	ctx, cancel := context.WithTimeout(req.Context(), tr.Timeout)
	defer cancel()
	req = req.WithContext(ctx)
	return tr.RoundTripper.RoundTrip(req)
}
