package transports

import (
	"net/http"
	"net/url"
)

// ProxyTransport uses the given proxy for all requests
type ProxyTransport struct {
	// The actual underlying transport we will use, shallow copy of the
	transport *http.Transport

	Proxy *url.URL
}

func (tr *ProxyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if tr.transport == nil {
		transport := http.DefaultTransport.(*http.Transport)
		tr.setTransport(transport)
	}
	return tr.transport.RoundTrip(req)
}

func (tr *ProxyTransport) proxy(req *http.Request) (*url.URL, error) {
	return tr.Proxy, nil
}

// shallow copy the given transport
func (tr *ProxyTransport) setTransport(transport *http.Transport) {
	*tr.transport = *transport
	tr.transport.Proxy = tr.proxy
}

func NewProxyTransport(transport *http.Transport) *ProxyTransport {
	tr := &ProxyTransport{}
	tr.setTransport(transport)
	return tr
}
