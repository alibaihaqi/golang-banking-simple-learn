package service

import (
	"github.com/alibaihaqi/banking/domain"
	"github.com/alibaihaqi/banking/dto"
	"github.com/alibaihaqi/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	var st string
	if status == "active" {
		st = "1"
	} else if status == "inactive" {
		st = "0"
	} else {
		st = ""
	}

	lc, err := s.repo.FindAll(st)
	if err != nil {
		return nil, err
	}

	sc := make([]dto.CustomerResponse, 0)
	for _, c := range lc {
		r := c.ToDto()
		sc = append(sc, r)
	}
	return sc, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	r := c.ToDto()
	return &r, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
