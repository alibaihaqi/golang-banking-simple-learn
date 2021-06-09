package domain

import (
	"database/sql"
	"fmt"
	"github.com/alibaihaqi/banking/errs"
	"github.com/alibaihaqi/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	r := d.client.QueryRow(customerSql, id)

	var c Customer
	err := r.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewNotFoundError("Customer is not found")
		} else {
			logger.Error("Error while scanning customer table:" + err.Error())
			return nil, errs.NewInternalSystemError("Unexpected Database Error!")
		}
	}

	return &c, nil
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	if status != "" {
		var s int
		if status == "active" {
			s = 1
		} else if status == "inactive" {
			s = 0
		}

		findAllSql = fmt.Sprintf("%s where status = %d", findAllSql, s)
	}

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		logger.Error("Error while querying customer table:" + err.Error())
		return nil, errs.NewInternalSystemError("Unexpected Database Error!")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scanning customer table:" + err.Error())
			return nil, errs.NewInternalSystemError(err.Error())
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:bHcbNUIe@6!&qiO@7xk7KAZuD^gDy&XZjMFJdsRTJtwts&mBrl@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
