package controllers

import (
	"AuthenticationModule/models"
	"AuthenticationModule/utils"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"gorm.io/gorm"
)

func SingUp(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.BadRequestResponse(utils.APIResponse{
			Context: ctx,
			LocalizeConfig: i18n.LocalizeConfig{
				MessageID: "ERROR_BINDING",
				TemplateData: map[string]interface{}{
					"Resouce": "user",
				},
			},
			Data: gin.H{},
		})
		return
	}

	if err := utils.Validate.StructPartial(user, "Email", "Password"); err != nil {
		utils.ValidatorErrorResponse(utils.APIResponse{
			Context: ctx,
			LocalizeConfig: i18n.LocalizeConfig{
				MessageID: "VALIDATION_ERROR",
				TemplateData: map[string]interface{}{
					"Resouce": "user",
				},
			},
			Data:           gin.H{},
			ValidatorError: err,
		})
		return
	}

	if err := user.Store(); err != nil {
		utils.InternalServerErrorResponse(utils.APIResponse{
			Context: ctx,
			LocalizeConfig: i18n.LocalizeConfig{
				MessageID: "STORE_ERROR",
				TemplateData: map[string]interface{}{
					"Resouce": "user",
				},
			},
			MessageError: err.Error(),
			Data:         gin.H{},
		})
		return
	}

	utils.OkResponse(utils.APIResponse{
		Context: ctx,
		LocalizeConfig: i18n.LocalizeConfig{
			MessageID: "SUCCESS_STORE",
			TemplateData: map[string]interface{}{
				"Resouce": "user",
			},
		},
		Data: user,
	})
}

func LogIn(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		utils.BadRequestResponse(utils.APIResponse{
			Context: ctx,
			LocalizeConfig: i18n.LocalizeConfig{
				MessageID: "ERROR_BINDING",
				TemplateData: map[string]interface{}{
					"Resouce": "user",
				},
			},
			Data: gin.H{},
		})
		return
	}

	if err := utils.Validate.StructPartial(user, "Email", "Password"); err != nil {
		utils.ValidatorErrorResponse(utils.APIResponse{
			Context: ctx,
			LocalizeConfig: i18n.LocalizeConfig{
				MessageID: "VALIDATION_ERROR",
				TemplateData: map[string]interface{}{
					"Resouce": "user",
				},
			},
			Data:           gin.H{},
			ValidatorError: err,
		})
		return
	}

	clientPassword := user.Password

	if err := user.ShowByCredentials(user.Email); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.BadRequestResponse(utils.APIResponse{
				Context: ctx,
				LocalizeConfig: i18n.LocalizeConfig{
					MessageID: "LOGIN_INVALID_EMAIL",
				},
				MessageError: err.Error(),
				Data:         gin.H{},
			})
			return
		}
		utils.InternalServerErrorResponse(utils.APIResponse{
			Context: ctx,
			LocalizeConfig: i18n.LocalizeConfig{
				MessageID: "DATABASE_ERROR",
			},
			MessageError: err.Error(),
			Data:         gin.H{},
		})
		return
	}

	if ok := utils.CheckPasswordHash(clientPassword, user.Password); !ok {
		utils.BadRequestResponse(utils.APIResponse{
			Context: ctx,
			LocalizeConfig: i18n.LocalizeConfig{
				MessageID: "LOGIN_INVALID_PASSWORD",
			},
			Data: gin.H{},
		})
		return
	}

	accessToken, err := utils.SignJWT(user.ID)
	if err != nil {
		utils.InternalServerErrorResponse(utils.APIResponse{
			Context: ctx,
			LocalizeConfig: i18n.LocalizeConfig{
				MessageID: "USER_DATA_ERROR",
			},
			MessageError: err.Error(),
			Data:         gin.H{},
		})
		return
	}

	refreshToken, err := utils.SignRefreshJWT(user.ID)
	if err != nil {
		utils.InternalServerErrorResponse(utils.APIResponse{
			Context: ctx,
			LocalizeConfig: i18n.LocalizeConfig{
				MessageID: "USER_DATA_ERROR",
			},
			MessageError: err.Error(),
			Data:         gin.H{},
		})
		return
	}

	utils.OkResponse(utils.APIResponse{
		Context: ctx,
		LocalizeConfig: i18n.LocalizeConfig{
			MessageID: "LOGIN_SUCCESS",
		},
		Data: gin.H{
			"user":         user,
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		},
	})
}

func LogOut(ctx *gin.Context) {

}
