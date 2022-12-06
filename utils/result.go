package utils

import (
	"net/http"
	resultdto "party/database/dto/result"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleSucces(c *gin.Context, data interface{}) {
	responseData := resultdto.Result{
		Status: "200",
		Message: "Success",
        Data:    data,
	}

	c.JSON(http.StatusOK, responseData)
}

func HandleFailed(c *gin.Context, status int, Message string) {
	responseData := resultdto.Result{
		Status: strconv.Itoa(status),
        Message: Message,
	}

	c.JSON(status, responseData)
}