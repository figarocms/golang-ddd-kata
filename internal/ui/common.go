package ui

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/GermainSIGETY/golang-ddd-kata/internal/domain/todo/model"
	"github.com/gin-gonic/gin"
)

const contentType = "Content-Type"
const JSONContentType = "application/json"

type ErrorsArrayJsonResponse struct {
	ErrorJsonResponse []ErrorJsonResponse `json:"errors"`
}
type ErrorJsonResponse struct {
	Code    *string `json:"code,omitempty"`
	Field   *string `json:"field,omitempty"`
	Message *string `json:"message,omitempty"`
}

func answerBadRequest(context *gin.Context, message string) {
	context.Header(contentType, JSONContentType)
	context.String(http.StatusBadRequest, "{\"errors\":[{\"message\": \"%v\"}]}", message)
}

func answerError(context *gin.Context, err interface{}) {
	errs, ok := err.([]model.DomainError)
	if !ok {
		answerError500(context, err)
		return
	}
	answerUnprocessableEntity(context, errs)
}

func answerErrorFromError(context *gin.Context, err error) {
	if errors.Is(err, &model.DomainError{}) {
		// handle domain error

		idError := &model.IDInferiorToZeroError{}
		if errors.As(err, &idError) {
			context.JSON(http.StatusBadRequest, fmt.Sprintf("id is inferior to 0 : %d", idError.Id))
			return
		}

		context.JSON(http.StatusBadRequest, "wrong input")
	} else if errors.Is(err, &model.TechnicalError{}) {
		// handle technical error
		answerError500(context, err)
		return
	}

	answerError500(context, err)
}

func answerResourceNotFound(context *gin.Context, message string) {
	notFound := "NOT_FOUND"
	errorJson := ErrorJsonResponse{Code: &notFound, Message: &message}
	context.JSON(http.StatusNotFound, ErrorsArrayJsonResponse{[]ErrorJsonResponse{errorJson}})
}

func answerUnprocessableEntity(context *gin.Context, errs []model.DomainError) {
	jsonErrors := make([]ErrorJsonResponse, len(errs))
	for i, v := range errs {
		code := v.Code()
		field := v.Field()
		description := v.Description()
		jsonErrors[i] = ErrorJsonResponse{&(code), &field, &description}
	}
	response := ErrorsArrayJsonResponse{jsonErrors}
	context.JSON(http.StatusUnprocessableEntity, response)
}

func answerError500(context *gin.Context, err interface{}) {
	fmt.Printf("ERROR : %v", err)
	context.Header(contentType, JSONContentType)
	context.String(http.StatusInternalServerError, "{\"errors\":[{\"code\": \"INTERNAL_ERROR\"}]}")
}
