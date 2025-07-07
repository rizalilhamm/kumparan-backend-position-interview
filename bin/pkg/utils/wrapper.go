package utils

import (
	"net/http"

	httpError "kumparan-backend-position-interview/bin/pkg/http-error"

	"github.com/labstack/echo/v4"
)

// Result common output
type Result struct {
	Data     interface{}
	MetaData interface{}
	Error    interface{}
	Count    int64
}

type Reply struct {
	Data  interface{}
	Error error
	Count *int64
	Meta  *MetaData `json:"meta,omitempty"`
}

type ResultCount struct {
	Data     int64
	MetaData interface{}
	Error    interface{}
}

// BaseWrapperModel data structure
type BaseWrapperModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Meta    interface{} `json:"meta,omitempty"`
}

type MetaData struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	TotalPage int `json:"total_page"`
	TotalData int `json:"total_data"`
}

// Response function
func Response(data interface{}, message string, code int, c echo.Context) error {
	success := false

	if code < http.StatusBadRequest {
		success = true
	}

	result := BaseWrapperModel{
		Success: success,
		Data:    data,
		Message: message,
		Code:    code,
	}

	return c.JSON(code, result)
}

// PaginationResponse function
func PaginationResponse(data interface{}, meta interface{}, message string, code int, c echo.Context) error {
	success := false

	if code < http.StatusBadRequest {
		success = true
	}

	result := BaseWrapperModel{
		Success: success,
		Data:    data,
		Meta:    meta,
		Message: message,
		Code:    code,
	}

	return c.JSON(code, result)
}

func Send(reply Reply, message string, code int, c echo.Context) error {
	success := false

	if code < http.StatusBadRequest {
		success = true
	}

	result := BaseWrapperModel{
		Success: success,
		Data:    reply.Data,
		Message: message,
		Code:    code,
		Meta:    reply.Meta,
	}

	return c.JSON(code, result)
}

// ResponseError function
func ResponseError(err interface{}, c echo.Context) error {
	errObj := getErrorStatusCode(err)
	result := BaseWrapperModel{
		Success: false,
		Data:    errObj.Data,
		Message: errObj.Message,
		Code:    errObj.Code,
	}

	return c.JSON(errObj.ResponseCode, result)
}

func getErrorStatusCode(err interface{}) httpError.CommonError {
	errData := httpError.CommonError{}

	switch obj := err.(type) {
	case httpError.BadRequest:
		errData.ResponseCode = http.StatusBadRequest
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.Unauthorized:
		errData.ResponseCode = http.StatusUnauthorized
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.Conflict:
		errData.ResponseCode = http.StatusConflict
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.NotFound:
		errData.ResponseCode = http.StatusNotFound
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.InternalServerError:
		errData.ResponseCode = http.StatusInternalServerError
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.TooManyRequest:
		errData.ResponseCode = http.StatusTooManyRequests
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.UnprocessableEntity:
		errData.ResponseCode = http.StatusUnprocessableEntity
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	case httpError.Forbidden:
		errData.ResponseCode = http.StatusForbidden
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		return errData
	default:
		errData.Code = http.StatusConflict
		return errData
	}
}
