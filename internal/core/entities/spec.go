package entities

import (
	"crypto/sha256"
	"encoding/json"
)

type Spec struct {
	Filename string
	SpecName string
	Hash     [32]byte
}

func CreateSpec(request Request, filename string) (Spec, error) {
	jsonData, err := json.Marshal(request)
	if err != nil {
		return Spec{}, err
	}

	shaSum := sha256.Sum256(jsonData)
	spec := Spec{
		Filename: filename,
		SpecName: request.Name,
		Hash:     shaSum,
	}
	return spec, nil
}
