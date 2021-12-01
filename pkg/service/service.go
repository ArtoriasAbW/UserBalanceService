package service

import (
	balance "UserBalanceService"
	"UserBalanceService/pkg/repository"
)

type Balance interface {
	GetUser(user balance.User) (balance.User, error)
	IncreaseBalance(operation balance.Operation) (uint64, error)
	DecreaseBalance(operation balance.Operation) (uint64, error)
	MakeTransfer(operation balance.TransferOperation) (uint64, uint64, error)
}

type Service struct {
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Balance: NewBalanceService(repos.Balance),
	}
}