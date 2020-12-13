package account

import "github.com/julioc98/robin/internal/app/entity"

// Repository Account interface
type Repository interface {
	Create(account *entity.Account) (int, error)
	Get(id int) (*entity.Account, error)
	GetByEmailAndPassword(email, password string) (*entity.Account, error)
}

// Service Account interface
type Service interface {
	Create(account *entity.Account) (int, error)
	Get(id int) (*entity.Account, error)
	GetByEmailAndPassword(email, password string) (*entity.Account, error)
}
