package common

// StandardResponse is the JSON response format used by our API
type StandardResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}
