package service

import (
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/configuration/rest_err"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
