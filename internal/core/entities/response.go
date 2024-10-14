package entities

type Response struct {
	StatusCode int    `json:"status_code"`
	Body       string `json:"body"`
}
