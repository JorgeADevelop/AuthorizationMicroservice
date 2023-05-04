package middlewares

import (
	"AuthenticationModule/utils"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func LocalizerMiddleware(bundle *i18n.Bundle) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lang := ctx.GetHeader("Accept-Language")
		if lang == "" {
			localizer := i18n.NewLocalizer(bundle, language.English.String())
			ctx.Set("localizer", localizer)
			ctx.Next()
			return
		}
		tag, err := language.Parse(lang)
		if err != nil {
			utils.OkResponse(
				utils.APIResponse{
					Context: ctx,
					LocalizeConfig: i18n.LocalizeConfig{
						MessageID: "INVALID_LANGUAGE",
					},
					Data: gin.H{},
				})
			return
		}
		localizer := i18n.NewLocalizer(bundle, tag.String())
		ctx.Set("localizer", localizer)
		ctx.Next()
	}
}
