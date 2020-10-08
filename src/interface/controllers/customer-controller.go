package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/vidu171/clean-architecture-go/domain"

	"github.com/vidu171/clean-architecture-go/usecases"
)

type CustomerController struct {
	customerInteractor usecases.CustomerInteractor
}

func NewCustomerController(customerInteractor usecases.CustomerInteractor) *CustomerController {
	return &CustomerController{customerInteractor}
}

func (controller *CustomerController) Add(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var customer domain.Customer
	err := json.NewDecoder(req.Body).Decode(&customer)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: "Invalid Payload"})
		return
	}
	err2 := controller.customerInteractor.CreateCustomer(customer)
	if err2 != nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(ErrorResponse{Message: err2.Error()})
		return
	}
	res.WriteHeader(http.StatusOK)
}
