package router

import (
	"golang-crud-gin/controller"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(
	tagsController *controller.TagsController,
	userController *controller.UserController,
	authController *controller.AuthController,
	mediaController *controller.MediaController,
	giftController *controller.GiftController,
	reedemController *controller.ReedemController,
	rateController *controller.RateController,
) *gin.Engine {

	router := gin.Default()

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")

	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", tagsController.FindAll)
	tagsRouter.GET("/:tagId", tagsController.FindById)
	tagsRouter.POST("", tagsController.Create)
	tagsRouter.PATCH("/:tagId", tagsController.Update)
	tagsRouter.DELETE("/:tagId", tagsController.Delete)

	userRouter := baseRouter.Group("/user")
	userRouter.GET("", userController.FindAll)
	userRouter.GET("/:userId", userController.FindById)
	userRouter.GET("/findemail/:email", userController.FindByEmail)
	userRouter.POST("", userController.Create)
	userRouter.PATCH("/:userId", userController.Update)
	userRouter.DELETE("/:userId", userController.Delete)

	authRouter := baseRouter.Group("/auth")
	authRouter.POST("", authController.Login)
	authRouter.POST("/signup", authController.SignUp)

	mediaRouter := baseRouter.Group("/upload")
	mediaRouter.POST("/file", mediaController.FileUpload)
	mediaRouter.POST("/remote", mediaController.RemoteUpload)

	giftRouter := baseRouter.Group("/gift")
	giftRouter.GET("", giftController.FindAll)
	giftRouter.GET("/:giftId", giftController.FindById)
	giftRouter.POST("", giftController.Create)
	giftRouter.PATCH("/:giftId", giftController.Update)
	giftRouter.DELETE("/:giftId", giftController.Delete)

	reedemRouter := baseRouter.Group("/reedem")
	reedemRouter.GET("", reedemController.FindAll)
	reedemRouter.GET("/:reedemId", reedemController.FindById)
	reedemRouter.POST("", reedemController.Create)
	reedemRouter.PATCH("/:reedemId", reedemController.Update)
	reedemRouter.DELETE("/:giftId", reedemController.Delete)

	rateRouter := baseRouter.Group("/rate")
	rateRouter.GET("", rateController.FindAll)
	rateRouter.GET("/:rateId", rateController.FindById)
	rateRouter.POST("", rateController.Create)
	rateRouter.PATCH("/:rateId", rateController.Update)
	rateRouter.DELETE("/:giftId", rateController.Delete)

	return router
}
