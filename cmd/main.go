package main

import (
	"golang-gin/application/config"
	"golang-gin/pkg/controller"
	_ "golang-gin/pkg/docs"
	"golang-gin/pkg/helper"
	"golang-gin/pkg/model"
	"golang-gin/pkg/repository"
	"golang-gin/pkg/router"
	"golang-gin/pkg/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	// Database
	log.Info().Msg("Started Server!")

	db := config.DatabaseConnection()
	if db != nil {
		log.Info().Msg("errors database")
	}

	// rd := config.RedisConnect()

	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})
	db.Table("users").AutoMigrate(&model.User{})
	db.Table("reedems").AutoMigrate(&model.Reedem{})
	db.Table("ratings").AutoMigrate(&model.Rating{})

	// db.Model(&model.Reedem{}).AddForeignKey("cust_id", "customers(cust_id)", "CASCADE", "CASCADE")
	// db.Model(&model.Rating{}).AddForeignKey("cust_id", "customers(cust_id)", "CASCADE", "CASCADE")

	// Repository
	tagsRepository := repository.NewTagsREpositoryImpl(db)
	userRepository := repository.NewUserREpositoryImpl(db)
	giftRepository := repository.NewGiftREpositoryImpl(db)
	reedemRepository := repository.NewReedemREpositoryImpl(db)
	rateRepository := repository.NewRateREpositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	userService := service.NewUserServiceImpl(userRepository, validate)
	mediaService := service.NewMediaServiceImpl(validate)
	giftService := service.NewGiftServiceImpl(giftRepository, validate)
	reedemService := service.NewReedemServiceImpl(reedemRepository, validate)
	rateService := service.NewRateServiceImpl(rateRepository, validate)

	// Controller
	tagsController := controller.NewTagsController(tagsService)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthController(userService)
	mediaController := controller.NewMediaController(mediaService)
	giftController := controller.NewGiftsController(giftService, mediaService)
	reedemController := controller.NewReedemController(reedemService, giftService)
	rateController := controller.NewRateController(rateService)

	// Router
	routes := router.NewRouter(
		tagsController,
		userController,
		authController,
		mediaController,
		giftController,
		reedemController,
		rateController,
	)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
