package controller

import (
	"golang-crud-gin/data/response"
	"golang-crud-gin/model"
	"golang-crud-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaController struct {
	mediaService service.MediaService
}

func NewMediaController(service service.MediaService) *MediaController {
	return &MediaController{
		mediaService: service,
	}
}

func (controller *MediaController) FileUpload(ctx *gin.Context) {
	formFile, _, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			response.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       map[string]interface{}{"data": "Select a file to upload"},
			})
		return
	}

	uploadUrl, err := controller.mediaService.FileUpload(model.File{File: formFile})
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			response.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       map[string]interface{}{"data": "Error uploading file"},
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		response.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       map[string]interface{}{"data": uploadUrl},
		})
}

func (controller *MediaController) RemoteUpload(ctx *gin.Context) {
	var url model.Url
	if err := ctx.BindJSON(&url); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			response.MediaDto{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Data:       map[string]interface{}{"data": err.Error()},
			})
		return
	}

	uploadUrl, err := controller.mediaService.RemoteUpload(url)

	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			response.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       map[string]interface{}{"data": "Error uploading file"},
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		response.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       map[string]interface{}{"data": uploadUrl},
		})
}
