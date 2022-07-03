package account

import (
	MODEL_ACCOUNT "go-user/model/account"
	"go-user/presenter"
)

type IAccountService interface {
	GetUserByUserID(userID string) (MODEL_ACCOUNT.User, *presenter.Response)
	GetUserService() ([]MODEL_ACCOUNT.User, *presenter.Response)
}
