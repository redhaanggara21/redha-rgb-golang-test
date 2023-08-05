package service

import (
	"golang-crud-gin/model"
)

type MediaService interface {
	FileUpload(file model.File) (string, error)
	RemoteUpload(url model.Url) (string, error)
}
