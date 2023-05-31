package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommonResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GinRespSuccessData(c *gin.Context, data interface{}) {
	t := CommonResponse{
		Code: http.StatusOK,
		Msg:  "success",
		Data: data,
	}
	GinRespJson(c, http.StatusOK, t)
}

func GinRespSuccess(c *gin.Context) {
	t := CommonResponse{
		Code: http.StatusOK,
		Msg:  "success",
	}
	GinRespJson(c, http.StatusOK, t)
}

func GinRespJson(c *gin.Context, status int, v interface{}) {
	c.JSON(status, v)
}
