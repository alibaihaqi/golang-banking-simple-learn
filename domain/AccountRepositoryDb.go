package domain

import (
	"database/sql"
	"github.com/alibaihaqi/banking/errs"
	"github.com/alibaihaqi/banking/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) GetBalance(aId string) (*Account, *errs.AppError) {
	customerSql := "select customer_id, opening_date, account_type, amount, status from accounts where account_id = ?"

	var a Account
	err := d.client.Get(&a, customerSql, aId)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewNotFoundError("Customer is not found")
		} else {
			logger.Error("Error while scanning customer table:" + err.Error())
			return nil, errs.NewInternalSystemError("Unexpected Database Error!")
		}
	}

	return &a, nil
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account" + err.Error())
		return nil, errs.NewInternalSystemError("Unexpected error from database!")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id account" + err.Error())
		return nil, errs.NewInternalSystemError("Unexpected error from database!")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func (d AccountRepositoryDb) UpdateAccount(tr Transaction) (*Account, *errs.AppError) {
	acc, err := d.GetBalance(tr.AccountId)
	if err != nil {
		return nil, err
	}

	if tr.TransactionType == "deposit" {
		acc.Amount = acc.Amount + tr.Amount
	} else if tr.TransactionType == "withdrawal" {
		acc.Amount = acc.Amount - tr.Amount
	}

	sqlUpdate := "UPDATE accounts set amount = ? where account_id = ?"
	_, err2 := d.client.Exec(sqlUpdate, acc.Amount, tr.AccountId)

	if err2 != nil {

		logger.Error("Error while creating new account" + err2.Error())
		return nil, errs.NewInternalSystemError("Unexpected error from database!")
	}

	return acc, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
