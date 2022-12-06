package controllers

import (
	"net/http"
	"party/person/repositories"
	"party/utils"
	jwtToken "party/utils/jwt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type handlerProfile struct {
	ProfileRepositories repositories.ProfileRepositories
}

func HandlerProfile(ProfileRepositories repositories.ProfileRepositories) *handlerProfile {
	return &handlerProfile{ProfileRepositories}
}

var c *gin.Engine

func (h *handlerProfile) GetProfile(ctx *gin.Context) {
	// ctx.Writer.Header().Set("Content-Type", "application/json")
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		utils.HandleFailed(ctx, http.StatusBadRequest, "unauthorized")
		return
	}

	token = strings.Split(token, " ")[1]
	_, err := jwtToken.DecodeToken(token)
	if err != nil {
		utils.HandleFailed(ctx, http.StatusUnauthorized, "unauthorized")
		return
	}


	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		utils.HandleFailed(ctx, http.StatusBadRequest, err.Error())
		return
	}

	profile, err := h.ProfileRepositories.GetProfile(id)
	if err != nil {
		utils.HandleFailed(ctx, http.StatusNotFound, "profile not found")
		return
	}

	utils.HandleSucces(ctx, profile)

}
