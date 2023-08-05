package controller

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
)

type GiftController struct {
	giftService  service.GiftService
	mediaService service.MediaService
}

func NewGiftsController(service service.GiftService, media service.MediaService) *GiftController {
	return &GiftController{
		giftService:  service,
		mediaService: media,
	}
}

// CreateTags		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body request.CreateTagsRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (controller *GiftController) Create(ctx *gin.Context) {
	log.Info().Msg("create gift")

	image, _, err := ctx.Request.FormFile("image")
	// formFile, _ := ctx.MultipartForm()
	helper.ErrorPanic(err)

	uploadUrl, err := controller.mediaService.FileUpload(model.File{File: image})

	if err != nil {
		webResponse := response.Response{
			Code:   http.StatusInternalServerError,
			Status: "error",
			Data:   uploadUrl,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, webResponse)
		return
	}

	// name := formFile.Value["name"]
	// description := formFile.Value["description"]
	// stock := cast.ToInt(formFile.Value["stock"])

	createGiftRequest := request.CreateGiftRequest{
		Name:        ctx.Request.FormValue("name"),
		Description: ctx.Request.FormValue("description"),
		Stock:       cast.ToInt(ctx.Request.FormValue("stock")),
		Image:       uploadUrl,
	}

	controller.giftService.Create(createGiftRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   createGiftRequest,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// UpdateTags		godoc
// @Summary			Update tags
// @Description		Update tags data.
// @Param			tagId path string true "update tags by id"
// @Param			tags body request.CreateTagsRequest true  "Update tags"
// @Tags			tags
// @Produce			application/json
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagId} [patch]
func (controller *GiftController) Update(ctx *gin.Context) {
	log.Info().Msg("update gift")
	updateGiftRequest := request.UpdateGiftRequest{}
	err := ctx.ShouldBindJSON(&updateGiftRequest)
	helper.ErrorPanic(err)

	giftId := ctx.Param("giftId")
	id, err := strconv.Atoi(giftId)
	helper.ErrorPanic(err)
	updateGiftRequest.Id = id

	controller.giftService.Update(updateGiftRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// DeleteTags		godoc
// @Summary			Delete tags
// @Description		Remove tags data by id.
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags/{tagID} [delete]
func (controller *GiftController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete gift")
	giftId := ctx.Param("giftId")
	id, err := strconv.Atoi(giftId)
	helper.ErrorPanic(err)
	controller.giftService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindByIdTags 		godoc
// @Summary				Get Single tags by id.
// @Param				tagId path string true "update tags by id"
// @Description			Return the tahs whoes tagId valu mathes id.
// @Produce				application/json
// @Tags				tags
// @Success				200 {object} response.Response{}
// @Router				/tags/{tagId} [get]
func (controller *GiftController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid gift")
	giftId := ctx.Param("giftId")
	id, err := strconv.Atoi(giftId)
	helper.ErrorPanic(err)

	giftResponse := controller.giftService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   giftResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// FindAllTags 		godoc
// @Summary			Get All tags.
// @Description		Return list of tags.
// @Tags			tags
// @Success			200 {obejct} response.Response{}
// @Router			/tags [get]
func (controller *GiftController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	giftResponse := controller.giftService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   giftResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
