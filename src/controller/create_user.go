package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/configuration/logger"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/configuration/validation"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/controller/model/request"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/model"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err:= c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)
			
		c.JSON(errRest.Code, errRest)
		return 
	}

	domain := model.NewUserDomain(
		userRequest.Email, 
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,
	))
}