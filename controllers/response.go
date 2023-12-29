package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理各种响应的包装函数

/*
{
    "code": 10001   错误码
    "msg": "xxxx"   错误的提示信息
    "data": {}      存放数据
}
*/

type ResponseData struct {
	Code    ResponseCode `json:"code"`
	Message interface{}  `json:"message"`
	Data    interface{}  `json:"data"`
}

func ResponseError(c *gin.Context, code ResponseCode) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	})
}

func ResponseErrorWithMsg(c *gin.Context, code ResponseCode, msg interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: msg,
		Data:    nil,
	})
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code:    CODESUCCESS,
		Message: CODESUCCESS.Msg(),
		Data:    data,
	})
}
