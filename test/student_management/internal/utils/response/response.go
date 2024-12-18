package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func WriteJson[T any](w http.ResponseWriter, status int, data T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error) ErrorResponse {
	return ErrorResponse{
		Status:  StatusError,
		Message: err.Error(),
	}
}

func ValidationError[T any](data T, errs validator.ValidationErrors) ErrorResponse {
	errMsg := make(map[string]string)
	structType := reflect.TypeOf(data)
	for _, err := range errs {
		field, fieldExists := structType.FieldByName(err.Field())
		var fieldName string
		if !fieldExists {
			fieldName = err.Field()
		} else {
			fieldName = field.Tag.Get("json")
		}
		switch err.ActualTag() {
		case "required":
			errMsg[fieldName] = fmt.Sprintf("%s is required", err.Field())
		case "email":
			errMsg[fieldName] = fmt.Sprintf("%s must be a valid email address", err.Field())
		case "unique":
			errMsg[fieldName] = fmt.Sprintf("%s must be an unique value", err.Field())
		default:
			errMsg[fieldName] = fmt.Sprintf("%s is invalid", err.Field())
		}
	}
	return ErrorResponse{
		Status:  StatusError,
		Message: "Validation failed",
		Error:   errMsg,
	}
}

func Success(data any, msg string) SuccessResponse {
	return SuccessResponse{
		Status:  StatusOK,
		Message: msg,
		Data:    data,
	}
}
