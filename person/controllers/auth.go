package controllers

import (
	"fmt"
	"log"
	"net/http"
	authdto "party/database/dto/auth"
	"party/database/models"
	"party/person/repositories"
	"party/utils"
	"party/utils/bcrypt"
	jwtToken "party/utils/jwt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt/v4"
)

type handlerAuth struct {
	AuthRepositories repositories.AuthRepositories
}


func HandlerAuth(AuthRepositories repositories.AuthRepositories) *handlerAuth {
	return &handlerAuth{AuthRepositories}
}


func (h *handlerAuth) Register(ctx *gin.Context) {

	request := new(authdto.RegisterRequest)
	if err := ctx.ShouldBindJSON(&request); err != nil {
		utils.HandleFailed(ctx, http.StatusBadRequest, "oops ShouldBindJSON error")
        return
	}

	validation := validator.New()
	err := validation.Struct(request)
    if err!= nil {
		utils.HandleFailed(ctx, http.StatusBadRequest, err.Error())
	}

	pasword, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		utils.HandleFailed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	profile := models.Profile{
		FullName: request.FullName,
		Email:    request.Email,
        Password: pasword,
	}

	claims := jwt.MapClaims{}
	claims["id"] = profile.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //token expired in 2 hours

	token, err := jwtToken.GenerateToken(&claims)
	if err != nil {
		log.Fatal("Failed To Generate")
		return
	}

	data, err := h.AuthRepositories.Register(profile)
	if err != nil {
		utils.HandleFailed(ctx, http.StatusInternalServerError, err.Error())
		fmt.Println(data)
	}

	RegisterResponse := authdto.Response{
		Token: token,
	}


	utils.HandleSucces(ctx, RegisterResponse)


}

func (h *handlerAuth) Login(ctx *gin.Context){

	request := new(authdto.LoginRequest)
	if err := ctx.ShouldBindJSON(&request); err!= nil {
		utils.HandleFailed(ctx, http.StatusBadRequest, "oops ShouldBindJSON error")
        return
	}

	Login := models.Profile{
		Email: request.Email,
		Password: request.Password,
	}

	Login, err := h.AuthRepositories.Login(Login.Email)
	if err!= nil {
		utils.HandleFailed(ctx, http.StatusBadRequest, "Email is not valid")
	}

	valid := bcrypt.CheckPasswordHash(request.Password, Login.Password)
	if !valid {
		utils.HandleFailed(ctx, http.StatusBadRequest, "Password is not correct")
	}

	claims := jwt.MapClaims{}
	claims["id"] = Login.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() //token expired in 2 hours

	token, err := jwtToken.GenerateToken(&claims)
	if err != nil {
		log.Fatal("Failed To Generate")
		return
	}

	loginResponse := authdto.Response{
		Token: token,
	}

	utils.HandleSucces(ctx, loginResponse)
}