package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/megaariii/go-employee-exam/helper"
	"github.com/megaariii/go-employee-exam/model/api"
)

type NotFoundError struct {
	Error string
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}
	InternalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	if exception, ok := err.(NotFoundError); ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := api.APIResponse{
			Code:   http.StatusNotFound,
			Status: "Not Found Error",
			Data:   exception.Error,
		}

		helper.EncodeResponse(writer, response)
		return true
	} else {
		return false
	}
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	if exception, ok := err.(validator.ValidationErrors); ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		response := api.APIResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		helper.EncodeResponse(writer, response)
		return true
	} else {
		return false
	}
}

func InternalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	response := api.APIResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	helper.EncodeResponse(writer, response)
}
