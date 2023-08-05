package controller

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"
	"golang-crud-gin/service"
	"golang-crud-gin/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type AuthController struct {
	userService service.UserService
}

func NewAuthController(service service.UserService) *AuthController {
	return &AuthController{
		userService: service,
	}
}

// CreateTags		godoc
// @Summary			Create user
// @Description		Save user data in Db.
// @Param			user body request.CreateUserRequest true "Create user"
// @Produce			application/json
// @Tags			user
// @Success			200 {object} respo
var jwtKey = []byte("my_secret_key")

func (controller *AuthController) Login(ctx *gin.Context) {
	log.Info().Msg("create user")

	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	helper.ErrorPanic(err)

	userResponse := controller.userService.FindByEmail(createUserRequest.Email)
	var exitUser = &userResponse
	if exitUser == nil {
		webResponse := response.Response{
			Code:   400,
			Status: "user doesn't exist",
			Data:   userResponse,
		}
		ctx.JSON(http.StatusOK, webResponse)
		return
	}

	errHash := utils.CompareHashPassword(createUserRequest.Password, exitUser.Password)

	if !errHash {
		ctx.JSON(400, gin.H{"Error": "Invalid password"})
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &model.Claims{
		Role: userResponse.Role,
		StandardClaims: jwt.StandardClaims{
			Subject:   userResponse.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		webResponse := response.Response{
			Code:   400,
			Status: "error",
			Data:   nil,
		}
		ctx.JSON(http.StatusOK, webResponse)
		return
	}

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userResponse,
		Token:  tokenString,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthController) SignUp(ctx *gin.Context) {
	log.Info().Msg("create user")

	createUserRequest := request.CreateUserRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)

	helper.ErrorPanic(err)
	log.Info().Msg(createUserRequest.Email)

	userResponse := controller.userService.FindByEmail(createUserRequest.Email)
	var exitUser = &userResponse

	if exitUser.Id != 0 {
		webResponse := response.Response{
			Code:   400,
			Status: "user already exist",
			Data:   exitUser,
		}
		ctx.JSON(http.StatusOK, webResponse)
		return
	}

	var errHash error
	createUserRequest.Password, errHash = utils.GenerateHashPassword(createUserRequest.Password)

	if errHash != nil {
		ctx.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	controller.userService.Create(createUserRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   createUserRequest,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AuthController) Home(ctx *gin.Context) {

	cookie, err := ctx.Cookie("token")

	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "user" && claims.Role != "admin" {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	ctx.JSON(200, gin.H{"success": "home page", "role": claims.Role})
}

func (controller *AuthController) Premium(ctx *gin.Context) {

	cookie, err := ctx.Cookie("token")

	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role != "admin" {
		ctx.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	ctx.JSON(200, gin.H{"success": "premium page", "role": claims.Role})
}

func (controller *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
	ctx.JSON(200, gin.H{"success": "user logged out"})
}

// ADDITIONAL FUNCTIONALITIES

func (controller *AuthController) ResetPassword(ctx *gin.Context) {

	updateUserRequest := request.UpdateUserRequest{}
	err := ctx.ShouldBindJSON(&updateUserRequest)
	helper.ErrorPanic(err)

	userResponse := controller.userService.FindByEmail(updateUserRequest.Email)
	var exitUser = &userResponse
	if exitUser == nil {
		webResponse := response.Response{
			Code:   400,
			Status: "user doesn't exist",
			Data:   userResponse,
		}
		ctx.JSON(http.StatusOK, webResponse)
		return
	}

	var errHash error
	updateUserRequest.Password, errHash = utils.GenerateHashPassword(updateUserRequest.Password)

	if errHash != nil {
		ctx.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	controller.userService.Update(updateUserRequest)

	ctx.JSON(200, gin.H{"success": "Password Updated"})
}
