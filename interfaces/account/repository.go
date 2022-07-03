package account

import (
	MODEL_ACCOUNT "go-user/model/account"
	"go-user/presenter"
)

type IAccountRepository interface {
	GetUser() ([]MODEL_ACCOUNT.User, *presenter.Response)
	GetUserByUserID(userID string) (MODEL_ACCOUNT.User, *presenter.Response)
}
