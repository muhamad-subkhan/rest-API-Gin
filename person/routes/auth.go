package routes

import (
	"party/config"
	"party/person/repositories"
	"party/person/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(ctx *gin.RouterGroup){
	Auth := repositories.RepositoriesAuth(config.Migration())

	h := controllers.HandlerAuth(Auth)

	ctx.POST("/register",h.Register)
	ctx.POST("/login",h.Login)
}