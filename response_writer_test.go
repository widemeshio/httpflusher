package httpflusher_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/widemeshio/httpflusher"
)

func TestResponseWriterNoFlusher(t *testing.T) {
	ow := newFakeNonFlushableWriter()
	fw := httpflusher.NewResponseWriter(ow)
	fw.WriteHeader(201)
	n, err := fw.Write([]byte{1, 2})
	require.NoError(t, err)
	require.Equal(t, 2, n)
	require.Equal(t, []byte{1, 2}, ow.writtenData)
	require.Equal(t, 201, ow.writtenStatusCode)
}

func TestResponseWriterWithFlusher(t *testing.T) {
	ow := newFakeFlushableWriter()
	fw := httpflusher.NewResponseWriter(ow)
	n, err := fw.Write([]byte{1, 2})
	require.NoError(t, err)
	require.Equal(t, 2, n)
	require.Equal(t, []byte{1, 2}, ow.writtenData)
	require.Equal(t, 1, ow.flushCount)
}

type fakeNonFlushableWriter struct {
	h                 http.Header
	writtenStatusCode int
	writtenData       []byte
}

func newFakeNonFlushableWriter() *fakeNonFlushableWriter {
	return &fakeNonFlushableWriter{
		h: http.Header{},
	}
}

func (w *fakeNonFlushableWriter) Header() http.Header {
	return w.h
}

func (w *fakeNonFlushableWriter) WriteHeader(statusCode int) {
	w.writtenStatusCode = statusCode
}

func (w *fakeNonFlushableWriter) Write(data []byte) (int, error) {
	w.writtenData = append(w.writtenData, data...)
	return len(data), nil
}

type fakeFlushableWriter struct {
	*fakeNonFlushableWriter
	flushCount int
}

func newFakeFlushableWriter() *fakeFlushableWriter {
	return &fakeFlushableWriter{
		newFakeNonFlushableWriter(),
		0,
	}
}

func (fw *fakeFlushableWriter) Flush() {
	fw.flushCount++
}
