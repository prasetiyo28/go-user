package account

import (
	repo "go-user/interfaces/account"
	MODEL_ACCOUNT "go-user/model/account"
	"go-user/presenter"
)

//Service interface
type Service struct {
	repo repo.IAccountRepository
}

//NewService create new use case
func NewService(repo repo.IAccountRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUser() ([]MODEL_ACCOUNT.User, *presenter.Response) {

	dataUser, errDetailUser := s.repo.GetUser()
	if errDetailUser != nil {
		return dataUser, &presenter.Response{
			Message: errDetailUser.Message,
		}
	}

	return dataUser, nil
}

func (s *Service) GetUserByUserID(userID string) (MODEL_ACCOUNT.User, *presenter.Response) {
	dataUser, errDetailUser := s.repo.GetUserByUserID(userID)
	if errDetailUser != nil {
		return dataUser, &presenter.Response{
			Message: errDetailUser.Message,
		}
	}

	return dataUser, nil
}
