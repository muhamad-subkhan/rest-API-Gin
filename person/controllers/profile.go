package controllers

import (
	"net/http"
	"party/person/repositories"
	"party/utils"
	"strconv"

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

	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
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
