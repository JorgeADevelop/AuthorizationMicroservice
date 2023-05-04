package main

import (
	"AuthenticationModule/models"
	"AuthenticationModule/routers"
	"AuthenticationModule/utils"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	utils.InitLogger()
	if err := godotenv.Load(); err != nil {
		utils.Logger.Fatal("error loading .env file", zap.String("err", err.Error()))
	}
	if err := models.InitMySQLDataBase(); err != nil {
		utils.Logger.Fatal("error initializing database", zap.String("err", err.Error()))
	}
	utils.InitValidator()
	bundle := utils.InitTranslator()
	routers.InitRouter(bundle)
}
