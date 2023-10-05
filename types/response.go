package types

// Response represents a json response.
type Response struct {
	Sum   int    `json:"sum,omitempty"`
	Error string `json:"error,omitempty"`
}
