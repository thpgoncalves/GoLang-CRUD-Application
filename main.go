package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/configuration/logger"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/controller"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/controller/routes"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/model/service"
)

func main() {
	logger.Info("Starting the application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router :=  gin.Default()
	
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}