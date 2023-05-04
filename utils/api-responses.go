package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type APIResponse struct {
	Context        *gin.Context
	LocalizeConfig i18n.LocalizeConfig
	MessageError   string
	Errors         gin.H
	ValidatorError error
	Data           interface{}
}

func OkResponse(apiResponse APIResponse) {
	localizer := apiResponse.Context.MustGet("localizer").(*i18n.Localizer)
	message := localizer.MustLocalize(&apiResponse.LocalizeConfig)

	apiResponse.Context.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"code":    http.StatusOK,
		"message": message,
		"data":    apiResponse.Data,
		"error":   nil,
		"errors":  gin.H{},
	})
}

func BadRequestResponse(apiResponse APIResponse) {
	localizer := apiResponse.Context.MustGet("localizer").(*i18n.Localizer)
	message := localizer.MustLocalize(&apiResponse.LocalizeConfig)

	apiResponse.Context.JSON(http.StatusBadRequest, gin.H{
		"status":  "Bad Request",
		"code":    http.StatusBadRequest,
		"message": message,
		"data":    gin.H{},
		"error":   apiResponse.MessageError,
		"errors":  apiResponse.Errors,
	})
}

func InternalServerErrorResponse(apiResponse APIResponse) {
	localizer := apiResponse.Context.MustGet("localizer").(*i18n.Localizer)
	message := localizer.MustLocalize(&apiResponse.LocalizeConfig)

	apiResponse.Context.JSON(http.StatusInternalServerError, gin.H{
		"status":  "Internal Server Error",
		"code":    http.StatusInternalServerError,
		"message": message,
		"data":    gin.H{},
		"error":   apiResponse.MessageError,
		"errors":  apiResponse.Errors,
	})
}

func ValidatorErrorResponse(apiResponse APIResponse) {
	localizer := apiResponse.Context.MustGet("localizer").(*i18n.Localizer)
	message := localizer.MustLocalize(&apiResponse.LocalizeConfig)

	validationErrors := apiResponse.ValidatorError.(validator.ValidationErrors)
	var errs []gin.H
	for _, validationError := range validationErrors {
		errs = append(errs, gin.H{
			"field":      validationError.StructField(),
			"validation": validationError.Error(),
		})
	}
	apiResponse.Context.JSON(http.StatusBadRequest, gin.H{
		"status":  "Bad Request",
		"code":    http.StatusBadRequest,
		"message": message,
		"data":    gin.H{},
		"error":   apiResponse.MessageError,
		"errors":  errs,
	})
}
