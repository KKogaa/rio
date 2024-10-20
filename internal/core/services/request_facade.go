package services

import (
	"github.com/KKogaa/rio/internal/core/entities"
)

type RequestFacade struct {
	fileService    *FileService
	requestService *RequestService
}

func NewRequestFacade(fileService *FileService,
	requestService *RequestService) *RequestFacade {

	return &RequestFacade{
		fileService:    fileService,
		requestService: requestService,
	}
}

// obtain the current file directory and file the file definition
// if not find the file definition based on the file path
// make the spec have an alias
// TODO: future handle chaining requests
func (r *RequestFacade) Send(filename string) (entities.Request, entities.Response, error) {
	request, err := r.fileService.SearchForSpec(filename)
	if err != nil {
		return entities.Request{}, entities.Response{}, err
	}

	response, err := r.requestService.MakeRequest(request)
	if err != nil {
		return entities.Request{}, entities.Response{}, err
	}

	return request, response, nil
}

func (r *RequestFacade) SendByPath(filepath string) (entities.Request, entities.Response, error) {
	request, err := r.fileService.GetRequestFromFile(filepath)
	if err != nil {
		return entities.Request{}, entities.Response{}, err
	}

	response, err := r.requestService.MakeRequest(request)
	if err != nil {
		return entities.Request{}, entities.Response{}, err
	}

	return request, response, nil
}
