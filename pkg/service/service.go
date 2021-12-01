package service

import "UserBalanceService/pkg/repository"

type Balance interface {

}

type Service struct {
	Balance
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}