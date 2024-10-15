package entities

type Response struct {
	StatusCode int                    `json:"status_code"`
	Body       map[string]interface{} `json:"body"`
}
