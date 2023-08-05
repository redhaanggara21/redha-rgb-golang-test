package service

import "golang-gin/pkg/model"

type MediaService interface {
	FileUpload(file model.File) (string, error)
	RemoteUpload(url model.Url) (string, error)
}
