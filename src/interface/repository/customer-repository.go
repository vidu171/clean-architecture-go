package repository

import (
	"github.com/vidu171/clean-architecture-go/domain"
)

type CustomerRepo struct {
	handler DBHandler
}

func NewCustomerRepo(handler DBHandler) CustomerRepo {
	return CustomerRepo{handler}
}

func (repo CustomerRepo) SaveCustomer(customer domain.Customer) error {
	err := repo.handler.SaveCustomer(customer)
	if err != nil {
		return err
	}
	return nil
}
