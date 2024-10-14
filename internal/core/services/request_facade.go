package services

import (
	"log"

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

func (r *RequestFacade) Send(filepath string) (entities.Request, entities.Response, error) {
	// obtain the current file directory and file the file definition
	// if not find the file definition based on the file path
	// make the spec have an alias

	// TODO: future handle chaining requests
	request, err := r.fileService.GetRequestFromFile(filepath)
	if err != nil {
		log.Println("entro 1")
		return entities.Request{}, entities.Response{}, err
	}
	log.Println(request)

	response, err := r.requestService.MakeRequest(request)
	if err != nil {
		log.Println("entro 2")
		return entities.Request{}, entities.Response{}, err
	}

	return request, response, nil
}

// func (r *RequestFacade) List(directoryPath string) {
//   // get the list of all the json files in the directory and list their names
//
// }
