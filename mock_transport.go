package transports

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// MockTransport silently discards requests without sending them and returns
// the Response set on it
// if Response is nil, a valid, empty HTTP response with no error is returned
type MockTransport struct {
	http.RoundTripper

	Response *http.Response
}

func (tr *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if tr.Response != nil {
		return tr.Response, nil
	} else {
		buf := &bytes.Buffer{}
		return &http.Response{
			Status:     "200 OK",
			StatusCode: 200,
			Proto:      "HTTP/1.0",
			ProtoMajor: 1,
			ProtoMinor: 0,
			Header:     make(http.Header),
			Body:       ioutil.NopCloser(buf),
			Request:    req,
		}, nil
	}
}
