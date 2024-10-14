package services

import (
	"encoding/json"
	"io"
	"os"

	"github.com/KKogaa/rio/internal/core/entities"
)

type FileService struct {
}

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
