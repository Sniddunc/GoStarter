package app

import (
	"log"
	"net/http"
)

// WrappedResponseWriter is a wrapped version of ResponseWriter which gives us access to the status code
type wrappedResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WrapResponseWriter wraps a ResponseWriter to give us access to the response's status code
func wrapResponseWriter(w http.ResponseWriter) wrappedResponseWriter {
	return wrappedResponseWriter{w, http.StatusOK}
}

// WriteHeader sets the status code of this response
func (wrw *wrappedResponseWriter) WriteHeader(statusCode int) {
	wrw.statusCode = statusCode
	wrw.ResponseWriter.WriteHeader(statusCode)
}

func logRequest(c *Context, roundTime int64) {
	log.Printf("| %d | %s %s | %dµs", c.Writer.statusCode, c.Request.Method, c.Request.URL.Path, roundTime)
}
