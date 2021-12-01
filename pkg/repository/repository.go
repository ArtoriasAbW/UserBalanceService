package repository

import (
	balance "UserBalanceService"
	"github.com/jmoiron/sqlx"
)

type Balance interface {
	GetUser(userId uint64) (balance.User, error)
	ModifyUser(user balance.User) (uint64, error)
	CreateUser(userId uint64) (uint64, error)
}

type Repository struct {
	Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Balance: NewBalancePostgres(db),
	}
}
