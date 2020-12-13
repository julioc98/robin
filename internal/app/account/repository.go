package account

import (
	gorm "github.com/jinzhu/gorm"
	"github.com/julioc98/robin/internal/app/entity"
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

// Create Account
func (r *postgresRepository) Create(a *entity.Account) (int, error) {
	if dbc := r.db.Create(a); dbc.Error != nil {
		return 0, dbc.Error
	}
	return a.ID, nil
}

// Get Account
func (r *postgresRepository) Get(id int) (*entity.Account, error) {
	var account entity.Account
	if dbc := r.db.First(&account, id); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &account, nil
}

func (r *postgresRepository) GetByEmailAndPassword(email, password string) (*entity.Account, error) {
	var account entity.Account
	if dbc := r.db.Where("email = ? AND password = ?", email, password).First(&account); dbc.Error != nil {
		return nil, dbc.Error
	}
	return &account, nil
}
