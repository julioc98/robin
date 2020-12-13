package transaction

import (
	"time"

	"github.com/julioc98/robin/internal/app/entity"
)

// Transaction entity
type Transaction struct {
	ID            int            `gorm:"primary_key" json:"id"`
	AccountID     string         `json:"account_id"`
	Account       entity.Account `json:"account"`
	OperationID   int            `json:"operation_type_id"`
	Operation     Operation      `json:"operation"`
	Amount        int            `json:"amount"`
	Category      string         `json:"category"`
	Subcategory   string         `json:"subcategory"`
	PurchaseID    string         `json:"purchase_id"`
	PsProductCode string         `json:"psProductCode"`
	PsProductName string         `json:"psProductName"`
	ForceAccept   bool           `json:"forceAccept"`
	CreatedAt     time.Time      `json:"createdAt"`
}

// Operation entity
type Operation struct {
	ID          int    `gorm:"primary_key"  json:"id"`
	Description string `json:"description"`
}
