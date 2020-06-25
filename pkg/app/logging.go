package app

import (
	"log"
	"net/http"

	"github.com/fatih/color"
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
	statusCode := getStatusColor(c.Writer.statusCode).Sprintf(" %d ", c.Writer.statusCode)
	requestMethod := color.New(color.Bold).Sprintf("%s", c.Request.Method)

	log.Printf("%s %s %s | %dµs", statusCode, requestMethod, c.Request.URL.Path, roundTime)
}

func getStatusColor(statusCode int) *color.Color {
	if statusCode >= 600 {
		return color.New(color.BgBlack)
	} else if statusCode >= 500 {
		return color.New(color.BgRed)
	} else if statusCode >= 400 {
		return color.New(color.BgRed)
	} else if statusCode >= 300 {
		return color.New(color.BgCyan)
	} else if statusCode >= 200 {
		return color.New(color.BgGreen)
	} else {
		return color.New(color.BgBlack)
	}
}
