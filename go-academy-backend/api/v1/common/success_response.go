package common

import "net/http"

type SuccessResponseCode string

const Success SuccessResponseCode = "success"

type SuccessResponse struct {
	Code    SuccessResponseCode `json:"code"`
	Message string              `json:"message"`
	Data    interface{}         `json:"data"`
}

func NewSuccessResponse(data interface{}) (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Code:    Success,
		Message: "Success",
		Data:    data,
	}
}

func NewSuccessResponseWithoutData() (int, SuccessResponse) {
	return http.StatusOK, SuccessResponse{
		Code:    Success,
		Message: "Success",
		Data:    nil,
	}
}
