package controller

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ReedemController struct {
	reedemService service.ReedemService
	giftService   service.GiftService
}

func NewReedemController(service service.ReedemService, giftService service.GiftService) *ReedemController {
	return &ReedemController{
		reedemService: service,
		giftService:   giftService,
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
func (controller *ReedemController) Create(ctx *gin.Context) {
	log.Info().Msg("create reedem")
	createReedemRequest := request.CreateReedemRequest{}
	// updateGiftRequest := request.UpdateGiftRequest{}

	err := ctx.ShouldBindJSON(&createReedemRequest)
	helper.ErrorPanic(err)

	controller.reedemService.Create(createReedemRequest)

	updateGiftRequest := controller.giftService.FindById(createReedemRequest.GiftID)

	if updateGiftRequest.Stock < createReedemRequest.Amount {
		webResponse := response.Response{
			Code:   http.StatusOK,
			Status: "Out of Stock",
			Data:   createReedemRequest,
		}
		ctx.Header("Content-Type", "application/json")
		ctx.JSON(http.StatusOK, webResponse)
		return
	} else {
		var amountLast = updateGiftRequest.Stock - createReedemRequest.Amount
		updateGiftRequest.Stock = amountLast
		controller.giftService.Update(updateGiftRequest)
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   createReedemRequest,
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
func (controller *ReedemController) Update(ctx *gin.Context) {
	log.Info().Msg("update gift")
	updateReedemRequest := request.UpdateReedemRequest{}
	err := ctx.ShouldBindJSON(&updateReedemRequest)
	helper.ErrorPanic(err)

	giftId := ctx.Param("giftId")
	id, err := strconv.Atoi(giftId)
	helper.ErrorPanic(err)
	updateReedemRequest.Id = id

	controller.reedemService.Update(updateReedemRequest)

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
func (controller *ReedemController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete gift")
	giftId := ctx.Param("giftId")
	id, err := strconv.Atoi(giftId)
	helper.ErrorPanic(err)
	controller.reedemService.Delete(id)

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
func (controller *ReedemController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid gift")
	giftId := ctx.Param("giftId")
	id, err := strconv.Atoi(giftId)
	helper.ErrorPanic(err)

	reedemResponse := controller.reedemService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   reedemResponse,
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
func (controller *ReedemController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	reedemResponse := controller.reedemService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   reedemResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
