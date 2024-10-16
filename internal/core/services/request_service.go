package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/KKogaa/rio/internal/core/entities"
)

// TODO: this shouldn't be a service but just a client
// different implementations using different requesters core request or gorrilla

// TODO: support adding cookies and other types of requests like websocket or graphql
// also add support for http 2
type RequestService struct {
}

func NewRequestService() *RequestService {
	return &RequestService{}
}

func (r *RequestService) MakeRequest(request entities.Request) (entities.Response, error) {
	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		return entities.Response{}, err
	}

	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}

	jsonBody, err := json.Marshal(request.Body)
	if err != nil {
		return entities.Response{}, err
	}

	req.Body = io.NopCloser(bytes.NewReader(jsonBody))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return entities.Response{}, err
	}

	defer resp.Body.Close()
	if resp.Body == nil {
		return entities.Response{StatusCode: resp.StatusCode}, nil
	}

	var responseBody map[string]interface{}
	switch resp.Header.Get("Content-Type") {
	case "application/json":
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&responseBody)
		if err != nil {
			return entities.Response{}, err
		}
	case "text/plain":
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		responseBody = map[string]interface{}{"body": buf.String()}
	// case "application/xml":
	// case "text/html":
	default:
	}

	return entities.Response{Body: responseBody,
		StatusCode: resp.StatusCode}, nil
}
