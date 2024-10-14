package services

import (
	"io"
	"net/http"

	"github.com/KKogaa/rio/internal/core/entities"
)

// TODO: this shouldn't be a service but just a client
// different implementations using different requesters core request or gorrilla
type RequestService struct {
}

func NewRequestService() *RequestService {
	return &RequestService{}
}

func (r *RequestService) MakeRequest(request entities.Request) (entities.Response, error) {
	// TODO: support adding cookies and other types of requests like websocket or graphql
	// also add support for http 2
	req, err := http.NewRequest(request.Method, request.Url, nil)
	if err != nil {
		return entities.Response{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return entities.Response{}, err
	}

	if resp.Request.Body == nil {
		return entities.Response{StatusCode: resp.StatusCode}, nil
	}

	defer resp.Request.Body.Close()
	body, err := io.ReadAll(resp.Request.Body)
	if err != nil {
		return entities.Response{}, err
	}

	//TODO: fix this entity, parse it to a map
	return entities.Response{Body: string(body),
		StatusCode: resp.StatusCode}, nil
}
