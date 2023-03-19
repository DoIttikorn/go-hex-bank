package repository

import "database/sql"

// not public
type customerRepositoryDB struct {
	db *sql.DB
}

func NewCustomerRepositoryDB(db *sql.DB) customerRepositoryDB {
	return customerRepositoryDB{
		db: db,
	}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	return []Customer{
		{
			Id:   1,
			Name: "Doit",
		},
		{
			Id:   2,
			Name: "Doit2",
		},
	}, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	return &Customer{
		Id:   1,
		Name: "Doit",
	}, nil
}
