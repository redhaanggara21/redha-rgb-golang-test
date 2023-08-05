package controller

import (
	"golang-gin/pkg/data/request"
	"golang-gin/pkg/data/response"
	"golang-gin/pkg/helper"
	"golang-gin/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type RateController struct {
	rateService service.RateService
}

func NewRateController(service service.RateService) *RateController {
	return &RateController{
		rateService: service,
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
func (controller *RateController) Create(ctx *gin.Context) {
	log.Info().Msg("create rated")
	createRateRequest := request.CreateRateRequest{}
	err := ctx.ShouldBindJSON(&createRateRequest)
	helper.ErrorPanic(err)

	controller.rateService.Create(createRateRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
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
func (controller *RateController) Update(ctx *gin.Context) {
	log.Info().Msg("update rate")
	updateRateRequest := request.UpdateRateRequest{}
	err := ctx.ShouldBindJSON(&updateRateRequest)
	helper.ErrorPanic(err)

	giftId := ctx.Param("giftId")
	id, err := strconv.Atoi(giftId)
	helper.ErrorPanic(err)
	updateRateRequest.Id = id

	controller.rateService.Update(updateRateRequest)

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
func (controller *RateController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete rate")
	rateId := ctx.Param("rateId")
	id, err := strconv.Atoi(rateId)
	helper.ErrorPanic(err)
	controller.rateService.Delete(id)

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
func (controller *RateController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid rate")
	rateId := ctx.Param("rateId")
	id, err := strconv.Atoi(rateId)
	helper.ErrorPanic(err)

	rateResponse := controller.rateService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   rateResponse,
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
func (controller *RateController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	rateResponse := controller.rateService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   rateResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
