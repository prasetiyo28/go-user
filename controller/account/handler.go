package account

import (
	"go-user/pkg/utils"
	"go-user/service/account"
)

var (
	ValidateRequest = utils.ValidateRequest
)

type AccountController struct {
	as *account.Service
}

func AccountControllerHandler(as *account.Service) *AccountController {
	handler := &AccountController{
		as: as,
	}
	return handler
}
