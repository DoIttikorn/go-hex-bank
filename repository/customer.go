package repository

// Port of CustomerRepository
type Customer struct {
	Id   int
	Name string
}

// Adapter interface
type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
}
