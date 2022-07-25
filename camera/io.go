package camera

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/go-playground/validator/v10"
)

func LoadFromJsonPath(path string) (*CameraDto, error) {
	jsonFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		return nil, err
	}

	return LoadFromJson(string(byteValue))
}

func LoadFromJson(jsonData string) (*CameraDto, error) {
	var cam CameraDto

	err := json.Unmarshal([]byte(jsonData), &cam)

	if err != nil {
		return nil, err
	}

	err = validator.New().Struct(cam)

	if err != nil {
		return nil, err
	}

	return &cam, nil
}
