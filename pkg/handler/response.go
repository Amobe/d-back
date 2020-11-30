package handler

// HelloResponse represents the response of hello API.
type HelloResponse struct {
	Index uint32 `json:"index"` // Index of request during the limit duration.
}
