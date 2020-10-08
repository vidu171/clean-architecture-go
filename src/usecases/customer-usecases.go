package usecases

import (
	"github.com/vidu171/clean-architecture-go/domain"
)

type CustomerInteractor struct {
	CustomerRepository domain.CustomerRepository
	Logger             Logger
}

func NewCustomerInteractor(CustomerRepository domain.CustomerRepository, Logger Logger) CustomerInteractor {
	return CustomerInteractor{CustomerRepository, Logger}
}

func (interactor *CustomerInteractor) CreateCustomer(customer domain.Customer) error {
	err := interactor.CustomerRepository.SaveCustomer(customer)
	if err != nil {
		interactor.Logger.Log(err.Error())
		return err
	}
	return nil
}
