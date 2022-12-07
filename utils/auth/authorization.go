package auth

import (
	"net/http"
	"party/utils"
	jwtToken "party/utils/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)
func Authorization() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		
		token := ctx.GetHeader("Authorization")
		// fmt.Printf(token, "ini token")
		if token == "" {
            ctx.AbortWithStatus(401)
			utils.HandleFailed(ctx, http.StatusBadRequest, "unauthorized")
			return
		}

		token = strings.Split(token, " ")[1]
		_, err := jwtToken.DecodeToken(token)
		// fmt.Println(token)
		// fmt.Println(claims,"decode")

		if err != nil {
            ctx.AbortWithStatus(401)
			utils.HandleFailed(ctx, http.StatusUnauthorized, "failed to decode token")
			return
		}

		// context := context.WithValue(ctx.Request.Context(), "token", claims)
		// r:= ctx.Request.WithContext(context)
		// r.Header.Set("Authorization", "Bearer "+token)

		// ctxx := context.WithValue(ctx.Request.Context(), "userinfo", claims)
		// r := ctx.Request.WithContext(ctxx)
		// http.HandlerFunc.ServeHTTP(ctx.Writer, r.WithContext(ctxx))
		
	}
}