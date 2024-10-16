package services

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/KKogaa/rio/internal/core/entities"
	"github.com/tidwall/pretty"
)

type FileService struct{}

func NewFileService() *FileService {
	return &FileService{}
}

func (f *FileService) GetRequestFromFile(filepath string) (entities.Request, error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		return entities.Request{}, err
	}

	fileContents, err := io.ReadAll(file)
	if err != nil {
		return entities.Request{}, err
	}

	var request entities.Request
	err = json.Unmarshal(fileContents, &request)
	if err != nil {
		return entities.Request{}, err
	}

	return request, nil
}

func (f *FileService) CreateRequestFile(filename string, name string,
	method string, url string) (string, error) {
	filename = filepath.Clean(filename)
	if filepath.Ext(filename) == "" {
		filename += ".json"
	}
	req := entities.Request{
		Name:    name,
		Method:  method,
		Body:    map[string]interface{}{},
		Headers: map[string]string{},
		Url:     url,
	}

	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		return filename, err
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return filename, err
	}

	prettyJsonData := pretty.Pretty(jsonData)
	_, err = file.Write(prettyJsonData)
	if err != nil {
		return filename, err
	}

	return file.Name(), nil
}

func (f *FileService) SearchForSpec(name string) (entities.Request, error) {
	// iterate overt the current working directory and find all the json files
	files, err := os.ReadDir(".")
	if err != nil {
		return entities.Request{}, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
			continue
		}

		request, err := f.GetRequestFromFile(file.Name())
		if err == nil && request.Name == name {
			return request, nil
		}
	}

	return entities.Request{}, errors.New("request name not found")
}
