package repository

import (
	balance "UserBalanceService"
	"github.com/jmoiron/sqlx"
)

type Balance interface {
	GetUser(user balance.User) (balance.User, error)

}

type Repository struct {
	Balance
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Balance: NewBalancePostgres(db),
	}
}
