This library will be very useful for you to log API Requests and Responses Including the duration time.

# Usage

See [test/test.go](test/test.go)

```go
package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bisratyalew/logger"
)

func main() {
	client := http.Client{
		Transport: httplogger.NewLoggedTransport(http.DefaultTransport, newLogger()),
	}

	client.Get("http://github.com")
}

type httpLogger struct {
	log *log.Logger
}

func newLogger() *httpLogger {
	return &httpLogger{
		log: log.New(os.Stderr, "log - ", log.LstdFlags),
	}
}

func (l *httpLogger) LogRequest(req *http.Request) {
	l.log.Printf(
		"Request %s %s",
		req.Method,
		req.URL.String(),
	)
}

func (l *httpLogger) LogResponse(req *http.Request, res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	l.log.Printf(
		"Response method=%s status=%d durationMs=%d %s",
		req.Method,
		res.StatusCode,
		duration,
		req.URL.String(),
	)
}

```

Output:

```
% go run example/example.go
log - 2018/07/04 12:59:39 Request GET http://github.com
log - 2018/07/04 12:59:39 Response method=GET status=302
durationMs=101 http://github.com
log - 2018/07/04 12:59:39 Request GET
http://github.com/bisratyalew
log - 2018/07/04 12:59:39 Response method=GET status=200
durationMs=138
http://github.com/bisratyalew
```

# LICENSE

The MIT License (MIT)

GOLANG HTTP LOGGER LIBRARY
Copyright (c) 2018 Bisrat Yalew (http://github.com/bisratyalew).

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
