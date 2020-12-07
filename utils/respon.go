package utils

import (
	"e-wallet-simple-api/model"
	"e-wallet-simple-api/constant"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleSuccess is function wrap in respon success
func HandleSuccessData(c *gin.Context, data interface{}) {
	var returnData = model.ResponWrapper{
		Success: true,
		Message: constant.Sucess,
		Data: data,
	}
	c.JSON(http.StatusOK, returnData)
}

func HandleSuccess(c *gin.Context, message string) {
	var data = model.ResponWrapperMessage{
		Success: true,
		Message: message,
	}
	c.JSON(http.StatusOK, data)
}

// HandleError is function wrap in respon failed
func HandleError(c *gin.Context, status int, message string) {
	var returnData = model.ResponWrapper{
		Success: false,
		Message: message,
	}
	c.JSON(status, returnData)
}