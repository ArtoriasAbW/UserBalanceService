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

func (r *BalancePostgres) GetUser(userId uint64) (balance.User, error) {
	var user balance.User
	query := fmt.Sprintf("select * from %s where id=$1", userTable)
	row := r.db.QueryRow(query, userId)
	if err := row.Scan(&user.Id, &user.Balance); err != nil {
		return user, err
	}
	return user, nil
}


func (r *BalancePostgres) ModifyUser(user balance.User) (uint64, error) {
	var userId uint64
	query := fmt.Sprintf("update %s set balance=$1 where id=$2 returning id", userTable)
	row := r.db.QueryRow(query, user.Balance, user.Id)
	if err := row.Scan(&userId); err != nil {
		return 0, err
	}
	return userId, nil
}