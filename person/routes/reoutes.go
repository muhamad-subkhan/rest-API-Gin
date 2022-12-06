package routes

import "github.com/gin-gonic/gin"

func Routes(ctx *gin.RouterGroup){
	AuthRoutes(ctx)
	ProfileRoutes(ctx)
}