package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 错误码
const (
	CodeSuccess       = 0
	CodeInvalid       = 400
	CodeUnauthorized  = 401
	CodeForbidden     = 403
	CodeNotFound      = 404
	CodeConflict      = 409
	CodeInternalError = 500
)

// Response 统一响应结构
type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

// AppError 业务错误
type AppError struct {
	Code int
	Msg  string
	Err  error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Msg + ": " + e.Err.Error()
	}
	return e.Msg
}

// NewError 创建业务错误
func NewError(code int, msg string) *AppError {
	return &AppError{Code: code, Msg: msg}
}

// Wrap 包装底层错误
func Wrap(code int, msg string, err error) *AppError {
	return &AppError{Code: code, Msg: msg, Err: err}
}

// OK 成功响应
func OK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{
		Code: CodeSuccess,
		Msg:  "success",
		Data: data,
	})
}

// Fail 失败响应
func Fail(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(httpStatusFromCode(appErr.Code), Response{
			Code: appErr.Code,
			Msg:  appErr.Msg,
		})
		return
	}
	c.JSON(http.StatusInternalServerError, Response{
		Code: CodeInternalError,
		Msg:  "internal server error",
	})
}

func httpStatusFromCode(code int) int {
	switch code {
	case CodeInvalid:
		return http.StatusBadRequest
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	case CodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
