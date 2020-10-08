package domain

type Customer struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type CustomerRepository interface {
	SaveCustomer(customer Customer) error
}
