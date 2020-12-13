package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julioc98/robin/internal/app/entity"
	"github.com/julioc98/robin/internal/app/transaction"
)

// IntegrationHandler Interface
type IntegrationHandler interface {
	Status(w http.ResponseWriter, r *http.Request)
	AddPurchase(w http.ResponseWriter, r *http.Request)
}

type integrationHandler struct {
	integrationService transaction.Service
}

// NewIntegrationHandler Create a new handler
func NewIntegrationHandler(integrationService transaction.Service) IntegrationHandler {
	return &integrationHandler{
		integrationService: integrationService,
	}
}

// Status Integration
func (ah *integrationHandler) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"message": "Operação realizada com sucesso.",
		"code": 0
	  }`))
}

// AddPurchase a Integration
func (ah *integrationHandler) AddPurchase(w http.ResponseWriter, r *http.Request) {
	var req entity.PurchaseAdd
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// type Transaction struct {
	// 	ID            int            `gorm:"primary_key" json:"id"`
	// 	AccountID     string         `json:"account_id"`
	// 	Account       entity.Account `json:"account"`
	// 	OperationID   int            `json:"operation_type_id"`
	// 	Operation     Operation      `json:"operation"`
	// 	Amount        int            `json:"amount"`
	// 	Category      string         `json:"category"`
	// 	Subcategory   string         `json:"subcategory"`
	// 	PurchaseID    string         `json:"purchase_id"`
	// 	PsProductCode string         `json:"psProductCode"`
	// 	PsProductName string         `json:"psProductName"`
	// 	ForceAccept   bool           `json:"forceAccept"`
	// 	CreatedAt     time.Time      `json:"createdAt"`
	// }

	trsct := &transaction.Transaction{
		AccountID:   req.AccountID,
		OperationID: 4,
		Amount:      req.TotalAmount.Amount,
		Category:    "ESSENTIAL",
		// Subcategory: req.Establishment.Mcc,
		PurchaseID:    req.PurchaseID,
		PsProductCode: req.PsProductCode,
		PsProductName: req.PsProductName,
		ForceAccept:   req.ForceAccept,
	}

	id, err := ah.integrationService.Create(trsct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	amount := 200000

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{
		"message": "Operação realizada com sucesso.",
		"code": 0,
		"authorization_id": %d,
		"balance": {
			"amount": %d,
			"currency_code": 986
		}
	}`, id, amount)))
}
