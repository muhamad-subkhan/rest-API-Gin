package routes

import (
	"party/config"
	"party/person/controllers"
	"party/person/repositories"
	"party/utils/auth"

	"github.com/gin-gonic/gin"
)


func ProfileRoutes(ctx *gin.RouterGroup) {
	Profile := repositories.RepositoriesProfile(config.Migration())
	h := controllers.HandlerProfile(Profile)


	// r := ctx.Use(auth.Authorization(c))

	// r.GET("/profile/:id", h.GetProfile)
	ctx.Use(auth.Authorization()).GET("/profile/:id", h.GetProfile)
	// ctx.GET("/profile/:id", h.GetProfile)

	// r.GET("/profile/:id", h.GetProfile)
	// ctx.GET("/profile/:id", h.GetProfile)


}

