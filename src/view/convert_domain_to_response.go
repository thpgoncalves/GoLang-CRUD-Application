package view

import (
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/controller/model/response"
	"github.com/thpgoncalves/GoLang-CRUD-Application/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID: "",
		Email: userDomain.GetEmail() ,
		Name: userDomain.GetName(),
		Age: userDomain.GetAge(), 
	}
}