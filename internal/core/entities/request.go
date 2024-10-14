package entities

type Request struct {
	Name    string                 `json:"name"`
	Method  string                 `json:"method"`
	Body    map[string]interface{} `json:"body"`
	Headers map[string]string      `json:"headers"`
	Url     string                 `json:"url"`
}
