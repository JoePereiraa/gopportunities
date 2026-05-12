package handler

import (
	"fmt"
	"net/http"

	"github.com/JoePereiraa/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func header(ctx *gin.Context) {
	ctx.Header("Content-type", "aplication/json")
}

func sendError(ctx *gin.Context, code int, msg string) {
	header(ctx)
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	header(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
