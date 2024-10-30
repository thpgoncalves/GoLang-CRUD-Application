package service

import (
	"fmt"

	"github.com/thpgoncalves/GoLang-CRUD-Application/src/configuration/logger"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/configuration/rest_err"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	
	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())
	
	return nil
}
