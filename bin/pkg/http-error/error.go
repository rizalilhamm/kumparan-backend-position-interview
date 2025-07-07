package error

import "net/http"

// CommonError struct
type CommonError struct {
	Code         int         `json:"code"`
	ResponseCode int         `json:"responseCode,omitempty"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

// BadRequest struct
type BadRequest struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewBadRequest
func NewBadRequest() BadRequest {
	errObj := BadRequest{}
	errObj.Message = "Bad Request"
	errObj.Code = http.StatusBadRequest

	return errObj
}

// NotFound struct
type NotFound struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewNotFound() NotFound {
	errObj := NotFound{}
	errObj.Message = "NotFound"
	errObj.Code = http.StatusNotFound

	return errObj
}

// Unauthorized struct
type Unauthorized struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewUnauthorized() Unauthorized {
	errObj := Unauthorized{}
	errObj.Message = "Unauthorized"
	errObj.Code = http.StatusUnauthorized

	return errObj
}

// Conflict struct
type Conflict struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewConflict() Conflict {
	errObj := Conflict{}
	errObj.Message = "Conflict"
	errObj.Code = http.StatusConflict

	return errObj
}

// InternalServerError struct
type InternalServerError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewInternalServerError() InternalServerError {
	errObj := InternalServerError{}
	errObj.Message = "Internal Server Error"
	errObj.Code = http.StatusInternalServerError

	return errObj
}

// Too many Requests
type TooManyRequest struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewTooManyRequest() TooManyRequest {
	errObj := TooManyRequest{}
	errObj.Message = "TooManyRequests"
	errObj.Code = http.StatusTooManyRequests

	return errObj
}

//Forbidden
type Forbidden struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewForbidden() Forbidden {
	errObj := Forbidden{}
	errObj.Message = "Request Forbidden"
	errObj.Code = http.StatusForbidden

	return errObj
}

//UnprocessableEntity
type UnprocessableEntity struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewUnprocessableEntity() UnprocessableEntity {
	errObj := UnprocessableEntity{}
	errObj.Message = "UnprocessableEntity"
	errObj.Code = http.StatusUnprocessableEntity

	return errObj
}
