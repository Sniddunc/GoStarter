package app

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// Context is passed into route handlers to provide an easy interface to the request object as well as response utils.
type Context struct {
	Request *http.Request
	Writer  wrappedResponseWriter
	DB      *sql.DB
}

// Query takes a query param key and returns it's value
func (c *Context) Query(key string, defaultValue string) string {
	result := c.Request.URL.Query()[key]

	if len(result) > 0 {
		return result[0]
	}

	return defaultValue
}

// JSON takes in a status code and some data and tries to parse the data into JSON before sending it to the client.
// A panic if JSON marshalling fails.
func (c *Context) JSON(statusCode int, data interface{}) {
	// Set JSON headers
	header := c.Writer.Header()
	if currentType := header["Content-Type"]; len(currentType) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}

	// Try to marshal the data into JSON
	dataBytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	// Write status code
	c.Writer.WriteHeader(statusCode)

	// Send data
	_, err = c.Writer.Write(dataBytes)
	if err != nil {
		panic(err)
	}
}

// SendUnauthorized informs the user that they are unauthorized
func (c *Context) SendUnauthorized() {
	// Send back 401 (Unauthorized)
	c.Writer.WriteHeader(http.StatusUnauthorized)
	c.Writer.Write([]byte("Unauthorized"))
}
