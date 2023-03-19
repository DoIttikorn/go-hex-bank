package repository

import "database/sql"

type userDBRepository struct {
	db *sql.DB
}

func NewUserRepositoryDB(db *sql.DB) userDBRepository {
	return userDBRepository{
		db: db,
	}
}

func (r userDBRepository) GetAll() ([]User, error) {
	return []User{
		{
			Id:   1,
			Name: "Doit",
		},
	}, nil
}

func (r userDBRepository) GetById(id int) (*User, error) {
	return &User{
		Id:   1,
		Name: "Doit",
	}, nil
}
