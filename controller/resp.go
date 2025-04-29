package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type respCode int

const (
	respCodeUnknown respCode = iota - 1
	respCodeOk
	respCodeParams
	respCodeDb
	respCodeNotFound
	respCodeAuth
)

func respOk(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": respCodeOk,
		"msg":  "ok",
		"data": data,
	})
}

func respError(c *gin.Context, code respCode, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  message,
	})
}
