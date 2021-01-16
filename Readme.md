# httpflusher
[![Test](https://github.com/widemeshio/httpflusher/workflows/Test/badge.svg)](https://github.com/widemeshio/httpflusher/actions?query=workflow%3ATest)

Here is httpflusher, to help you with your chunked responses and buffering of the HTTP protocol endeavors.

`http.ResponseWriter` `+` automatic `http.Flusher`.

## Sample

```go
package main

import (
	"net/http"
	"time"

	"github.com/widemeshio/httpflusher"
)

func main() {
	var chunkedEndpoint = func(w http.ResponseWriter, req *http.Request) {
		w = httpflusher.NewResponseWriter(w)
		w.WriteHeader(200)
		w.Write([]byte("log-line-1\n"))
		time.Sleep(time.Second)
		w.Write([]byte("log-line-2\n"))
	}
	http.HandleFunc("/stream-logs", chunkedEndpoint)
	http.ListenAndServe(":8090", nil)
}
```

Run:

```sh
$ curl --no-buffer localhost:8090/stream-logs
```
