package service

import (
	"golang-gin/pkg/helper"
	"golang-gin/pkg/model"

	"github.com/go-playground/validator/v10"
)

type media struct{}

var (
	validate = validator.New()
)

type MediaServiceImpl struct {
	Validate *validator.Validate
}

func NewMediaServiceImpl(validate *validator.Validate) MediaService {
	return &MediaServiceImpl{
		Validate: validate,
	}
}

func (media *MediaServiceImpl) FileUpload(file model.File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := helper.ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (media *MediaServiceImpl) RemoteUpload(url model.Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := helper.ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}
