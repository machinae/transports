package transports

import (
	"io"
	"net/http"
	"net/http/httputil"
	"os"
)

// LoggingTransport dumps all outbound requests in wire format
type LoggingTransport struct {
	http.RoundTripper

	// Writer to dump requests. If nil, defaults to os.Stdout
	Writer io.Writer

	// If true, dump only headers, not request body
	SkipBody bool
}

func (tr *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if tr.RoundTripper == nil {
		tr.RoundTripper = http.DefaultTransport
	}
	if tr.Writer == nil {
		tr.Writer = os.Stdout
	}
	rawReq, err := httputil.DumpRequestOut(req, !tr.SkipBody)
	if err != nil {
		return nil, err
	}
	tr.Writer.Write(rawReq)
	return tr.RoundTripper.RoundTrip(req)
}
