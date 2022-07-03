package repository

import (
	MODEL_ACCOUNT "go-user/model/account"
	"go-user/pkg"
	"go-user/presenter"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

//Account new repository
func NewRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

// func (repository *AccountRepository) Insert(model *account.Init) {
// 	context := context.Background()
// 	repository.rc.Client.Set(context, "account:"+model.Token, model.CustomerXID, 0)
// }

func (repository *AccountRepository) GetUser() ([]MODEL_ACCOUNT.User, *presenter.Response) {
	var users []MODEL_ACCOUNT.User
	data := repository.db.Find(&users)
	if data.Error != nil {
		return users, &presenter.Response{
			Message: pkg.ErrInternalServer.Error(),
		}
	}
	return users, nil

}

func (repository *AccountRepository) GetUserByUserID(userID string) (MODEL_ACCOUNT.User, *presenter.Response) {
	var user MODEL_ACCOUNT.User
	data := repository.db.First(&user, userID)
	if data.Error != nil {
		if data.Error.Error() == gorm.ErrRecordNotFound.Error() {
			return user, &presenter.Response{
				Message: pkg.ErrDataNotFound.Error(),
			}
		}
		return user, &presenter.Response{
			Message: pkg.ErrInternalServer.Error(),
		}
	}
	return user, nil

}
