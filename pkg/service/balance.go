package service

import (
	balance "UserBalanceService"
	"UserBalanceService/pkg/repository"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance)  *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) GetUser(user balance.User)  (balance.User, error){
	return s.repo.GetUser(user)
}