package db

import (
	"github.com/jinzhu/gorm"
	"github.com/julioc98/robin/internal/app/entity"
	"github.com/julioc98/robin/internal/app/transaction"
)

// Migrate migration BD
func Migrate(conn *gorm.DB) {
	// Migrate the schema
	conn.AutoMigrate(&entity.Account{}, &transaction.Operation{}, &transaction.Transaction{})

	// Create an Account
	conn.Create(&entity.Account{ID: "1", DocumentNumber: "54010024992"})

	// Create Operactions
	conn.Create(&transaction.Operation{ID: 1, Description: "CASH-IN"})
	conn.Create(&transaction.Operation{ID: 2, Description: "CHARGEBACK"})
	conn.Create(&transaction.Operation{ID: 3, Description: "CANCEL"})
	conn.Create(&transaction.Operation{ID: 4, Description: "CASH-OUT"})

	//Create a Transaction
	conn.Create(&transaction.Transaction{
		AccountID:   "1",
		OperationID: 1,
		Amount:      5000,
	})
}
