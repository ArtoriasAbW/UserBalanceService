package service

import (
	balance "UserBalanceService"
	"UserBalanceService/pkg/repository"
	"errors"
)

type BalanceService struct {
	repo repository.Balance
}


func NewBalanceService(repo repository.Balance)  *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) GetUser(user balance.User)  (balance.User, error){
	return s.repo.GetUser(user.Id)
}

func (s *BalanceService) IncreaseBalance(operation balance.Operation) (uint64, error) {
	user, err := s.repo.GetUser(operation.UserId)
	if err != nil {
		return 0, errors.New("no such user")
	}
	user.Balance += operation.Value
	return s.repo.ModifyUser(user)
}