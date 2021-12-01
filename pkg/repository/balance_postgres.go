package repository

import (
	balance "UserBalanceService"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type BalancePostgres struct {
	db *sqlx.DB
}

func NewBalancePostgres(db *sqlx.DB) *BalancePostgres {
	return &BalancePostgres{db: db}
}

func (r *BalancePostgres) GetUser(user balance.User) (balance.User, error) {
	query := fmt.Sprintf("select * from %s where id=$1", userTable)
	row := r.db.QueryRow(query, user.Id)
	if err := row.Scan(&user.Id, &user.Balance); err != nil {
		return user, err
	}
	return user, nil
}