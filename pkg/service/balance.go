package service

import (
	balance "UserBalanceService"
	"UserBalanceService/pkg/repository"
	"errors"
	"math"
)

type BalanceService struct {
	repo repository.Balance
}

func NewBalanceService(repo repository.Balance)  *BalanceService {
	return &BalanceService{repo: repo}
}

func (s *BalanceService) GetUser(user balance.User)  (balance.User, error) {
	rUser, err := s.repo.GetUser(user.Id)
	if err != nil {
		return user, err
	}
	rUser.Currency = user.Currency
	if rUser.Currency != "" {
		rate, err := GetRate(rUser.Currency)
		if err != nil {
			return balance.User{}, err
		}
		rUser.Balance = math.Round(rUser.Balance * rate * 1000) / 1000
	}
	return rUser, nil
}

func (s *BalanceService) IncreaseBalance(operation balance.Operation) (uint64, error) {
	user, err := s.repo.GetUser(operation.UserId)
	if err != nil {
		userId, err := s.repo.CreateUser(operation.UserId)
		if err != nil {
			return 0, err
		}
		user = balance.User{
			Id: userId,
			Balance: 0,
		}
	}
	user.Balance += operation.Value
	return s.repo.ModifyUser(user)
}

func (s *BalanceService) DecreaseBalance(operation balance.Operation) (uint64, error) {
	user, err := s.repo.GetUser(operation.UserId)
	if err != nil {
		return 0, errors.New("no such user")
	}
	if user.Balance < operation.Value {
		return 0, errors.New("user doesn't have enough money")
	}
	user.Balance -= operation.Value
	return s.repo.ModifyUser(user)
}

func (s *BalanceService) MakeTransfer(operation balance.TransferOperation) (uint64, uint64, error) {
	sender, senderErr := s.repo.GetUser(operation.SenderId)
	receiver, receiverErr := s.repo.GetUser(operation.ReceiverId)
	if senderErr != nil {
		return 0, 0, errors.New("no such sender")
	}
	if receiverErr != nil {
		receiverId, err := s.repo.CreateUser(operation.ReceiverId)
		if err != nil {
			return 0, 0, err
		}
		receiver = balance.User{
			Id: receiverId,
			Balance: 0,
		}
	}
	if sender.Balance < operation.Value {
		return 0, 0, errors.New("sender doesn't have enough money")
	}
	sender.Balance -= operation.Value
	receiver.Balance += operation.Value
	senderId, senderErr := s.repo.ModifyUser(sender)
	if senderErr != nil {
		return 0, 0, senderErr
	}
	receiverId, receiverErr := s.repo.ModifyUser(receiver)
	return senderId, receiverId, receiverErr
}