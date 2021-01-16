package httpflusher

import (
	"net/http"
)

// ResponseWriter is a HTTP response writer with automatic flushing when available
type ResponseWriter interface {
	http.ResponseWriter
	http.Flusher
}

// NewResponseWriter creates a new response writer
// with automatic flushing via http.Flusher when available by the given response writer
func NewResponseWriter(w http.ResponseWriter) ResponseWriter {
	f, _ := w.(http.Flusher)
	return &responseWriter{
		w,
		f,
	}
}

type responseWriter struct {
	http.ResponseWriter
	flusher http.Flusher
}

func (f *responseWriter) Write(p []byte) (n int, err error) {
	n, err = f.ResponseWriter.Write(p)
	f.Flush()
	return
}

func (f *responseWriter) Flush() {
	if flusher := f.flusher; flusher != nil {
		flusher.Flush()
	}
}
