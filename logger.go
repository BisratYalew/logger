// HTTP API Request Logger 

package logger

// Load Dependencies
import (
	"log"
	"net/http"
	"time"
)

type RoundLogged struct {
	rt  http.RoundTripper
	log HTTPLogger
}

func (c *RoundLogged) RoundTrip(request *http.Request) (*http.Response, error) {
	c.log.LogReq(request) // Log the Request
	st := time.Now() // Start time
	response, err := c.rt.RoundTrip(request)
	duration := time.Since(st)
	c.log.LogRes(request, response, err, duration)
	return response, err
}

func NewLoggedTransport(rt http.RoundTripper, log HTTPLogger) http.RoundTripper {
	return &RoundLogged{rt: rt, log: log}
}

// HTTPLogger defines the interface to log all http request and responses
type HTTPLogger interface {
	LogReq(*http.Request)
	LogRes(*http.Request, *http.Response, error, time.Duration)
}

type DefaultLogger struct {
}

//Log Requests
func (dl DefaultLogger) LogReq(*http.Request) {
}

// Log Response
func (dl DefaultLogger) LogRes(req *http.Request, res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	if err != nil {
		log.Printf("HTTP Request method=%s host=%s path=%s status=error durationMs=%d error=%q", req.Method, req.Host, req.URL.Path, duration, err.Error())
	} else {
		log.Printf("HTTP Request method=%s host=%s path=%s status=%d durationMs=%d", req.Method, req.Host, req.URL.Path, res.StatusCode, duration)
	}
}

var DefaultLoggedTransport = NewLoggedTransport(http.DefaultTransport, DefaultLogger{})
