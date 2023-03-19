package service

import (
	"log"

	"github.com/Doittikorn/bank/repository"
)

// don't another package can access this struct
type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{
		custRepo: custRepo,
	}
}

func (s customerService) GetCustomer() ([]CustomerResponse, error) {
	customer, err := s.custRepo.GetAll()
	if err != nil {
		return nil, err
	}

	custResponse := []CustomerResponse{}
	for _, v := range customer {
		custResponse = append(custResponse, CustomerResponse{
			Name: v.Name,
		})
	}
	return custResponse, nil
}

func (s customerService) GetCustomerById(int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(1)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custResponse := CustomerResponse{
		Name: customer.Name,
	}

	return &custResponse, nil
}
