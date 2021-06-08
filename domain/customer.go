package domain

import "github.com/alibaihaqi/banking/errs"

type Customer struct {
	Id          string `json:"id" xml:"id"`
	Name        string `json:"full_name" xml:"name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zip_code" xml:"zipcode"`
	DateOfBirth string `json:"date_of_birth" xml:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
