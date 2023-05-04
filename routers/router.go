package routers

import (
	"AuthenticationModule/controllers"
	"AuthenticationModule/middlewares"
	"AuthenticationModule/utils"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func InitRouter(bundle *i18n.Bundle) {
	router := gin.Default()

	apiv1 := router.Group("/api")
	{
		apiv1.Use(middlewares.LocalizerMiddleware(bundle))
		apiv1.GET("/", func(ctx *gin.Context) {
			utils.OkResponse(
				utils.APIResponse{
					Context: ctx,
					LocalizeConfig: i18n.LocalizeConfig{
						MessageID: "HEALTHY_CHECK",
					},
					Data: gin.H{},
				})
		})

		apiv1.POST("/sing-up", controllers.SingUp)
		apiv1.POST("/log-in", controllers.LogIn)
		apiv1.POST("/log-out", controllers.LogOut)
	}
	router.Run()
}
