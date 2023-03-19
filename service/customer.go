package service

type CustomerResponse struct {
	Name string `json:"name"`
}

type CustomerService interface {
	GetCustomer() ([]CustomerResponse, error)
	GetCustomerById(int) (*CustomerResponse, error)
}
