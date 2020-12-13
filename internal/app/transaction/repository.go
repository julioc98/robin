package transaction

import (
	"github.com/jinzhu/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

//NewPostgresRepository create new postgres repository
func NewPostgresRepository(db *gorm.DB) Repository {
	return &postgresRepository{
		db,
	}
}

// Create Transaction
func (r *postgresRepository) Create(a *Transaction) (int, error) {
	if dbc := r.db.Create(a); dbc.Error != nil {
		return 0, dbc.Error
	}
	return a.ID, nil
}

// Get Transaction
func (r *postgresRepository) Get(id int) (*Transaction, error) {
	var transaction Transaction

	if dbc := r.db.Set("gorm:auto_preload", true).First(&transaction, id); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &transaction, nil
}

// Get Transaction
func (r *postgresRepository) GetFounds(accountID, category string) (int, error) {
	type Result struct {
		Amount int
	}
	var result Result
	if dbc := r.db.Raw("SELECT SUM(amount) as amount FROM transactions WHERE account_id=? AND category = ?", accountID, category).Scan(&result); dbc.Error != nil {
		return 0, dbc.Error
	}
	return result.Amount, nil
}
