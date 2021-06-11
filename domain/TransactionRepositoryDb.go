package domain

import (
	"github.com/alibaihaqi/banking/errs"
	"github.com/alibaihaqi/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func (d TransactionRepositoryDb) NewTransaction(t Transaction) (*Transaction, *errs.AppError) {
	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while creating new transactions" + err.Error())
		return nil, errs.NewInternalSystemError("Unexpected error from database!")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id account" + err.Error())
		return nil, errs.NewInternalSystemError("Unexpected error from database!")
	}

	t.TransactionId = strconv.FormatInt(id, 10)

	return &t, nil
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient}
}
