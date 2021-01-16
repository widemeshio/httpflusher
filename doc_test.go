package httpflusher_test

import (
	"net/http"
	"time"

	"github.com/widemeshio/httpflusher"
)

func Example() {
	var chunkedEndpoint = func(w http.ResponseWriter, req *http.Request) {
		w = httpflusher.NewResponseWriter(w)
		w.WriteHeader(200)
		w.Write([]byte("log-line-1\n"))
		time.Sleep(time.Second)
		w.Write([]byte("log-line-2\n"))
	}
	http.HandleFunc("/stream-logs", chunkedEndpoint)
	http.ListenAndServe(":8090", nil)

	// $ curl --no-buffer localhost:8090/stream-logs
}
