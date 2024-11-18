package utils

import (
	"net/http"
	"zvaljean/proxy-subscribe-api/pkg/common/errno"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespStr(ctx *gin.Context, data string) {
	ctx.String(http.StatusOK, data)
}
func RespOk(ctx *gin.Context, data interface{}) {

	ctx.JSON(http.StatusOK, Response{
		Code: errno.OK.Code,
		Msg:  errno.OK.Msg,
		Data: data})
}

func RespBiz(ctx *gin.Context, bc *errno.BizCode) {

	ctx.JSON(http.StatusOK, Response{
		Code: bc.Code,
		Msg:  bc.Error(),
		Data: nil})
}

func RespErr(ctx *gin.Context, bc *errno.BizCode) {

	ctx.JSON(http.StatusInternalServerError, Response{
		Code: bc.Code,
		Msg:  bc.Error(),
		Data: nil})
}
