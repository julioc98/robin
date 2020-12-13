package processor

import "github.com/julioc98/robin/internal/app/entity"

// // Repository Account interface
// type Repository interface {
// 	Create(account *Account) (int, error)
// 	Get(id int) (*Account, error)
// }

// // Service Account interface
// type Service interface {
// 	Create(account *Account) (int, error)
// 	Get(id int) (*Account, error)
// }

type Gateway interface {
	CreateAccount(account *entity.Account) (string, error)
	CreateVirtualCard(accountID string) (*entity.VirtualCard, error)
}
