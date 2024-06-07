package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Response 统一响应格式
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// SendResponse 发送通用响应
func SendResponse(ctx echo.Context, code int, msg string, data interface{}) error {
	response := Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	return ctx.JSON(code, response)
}

// Success 返回成功响应
func Success(ctx echo.Context, msg string, data interface{}) error {
	return SendResponse(ctx, http.StatusOK, msg, data)
}

// Fail 返回错误响应
func Fail(ctx echo.Context, code int, msg string) error {
	return SendResponse(ctx, code, msg, nil)
}

func InternalServerError(ctx echo.Context, err error) error {
	return SendResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
}
