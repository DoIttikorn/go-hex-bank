package repository

import "database/sql"

type accountRepositoryDB struct {
	db *sql.DB
}

func NewAccountRepositoryDB(db *sql.DB) AccountRepository {
	return &accountRepositoryDB{db}
}

func (r *accountRepositoryDB) Create(account Account) (*Account, error) {
	query, err := r.db.Prepare("INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	result, err := query.Exec(account.CustomerID, account.OpeningDate, account.AccountType, account.Amount, account.Status)
	if err != nil {
		return nil, err
	}
	last, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	account.AccountID = int(last)

	return &account, nil
}

func (r *accountRepositoryDB) GetAll(customerID int) ([]Account, error) {
	query, err := r.db.Query("SELECT * FROM accounts WHERE customer_id = ?", customerID)
	if err != nil {
		return nil, err
	}

	accounts := []Account{}
	for query.Next() {
		var account Account
		query.Scan(&account.AccountID, &account.CustomerID, &account.OpeningDate, &account.AccountType, &account.Amount, &account.Status)
		accounts = append(accounts, account)
	}

	return accounts, nil
}
