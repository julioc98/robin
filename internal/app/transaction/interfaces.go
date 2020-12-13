package transaction

import "github.com/julioc98/robin/internal/app/entity"

// Repository Account interface
type Repository interface {
	GetFounds(accountID, category string) (int, error)
	Create(account *Transaction) (int, error)
	Get(id int) (*Transaction, error)
}

// Service Account interface
type Service interface {
	GetFounds(accountID string) ([]entity.Found, error)
	AddFunds(t *Transaction) (int, error)
	Create(account *Transaction) (int, error)
	Get(id int) (*Transaction, error)
	GetFoundByCategory(accountID, category string) (int, error)
}
