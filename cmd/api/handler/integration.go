package handler

import (
	"encoding/json"
	"fmt"
	"log"
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
		"code": 00
	  }`))
}

// AddPurchase a Integration
func (ah *integrationHandler) AddPurchase(w http.ResponseWriter, r *http.Request) {
	var req entity.PurchaseAdd
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("[ERROR]", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	amount, err := ah.integrationService.GetFoundByCategory(req.AccountID, trsct.Category)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("[AMOUNT]", amount)

	if amount < trsct.Amount {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := ah.integrationService.Create(trsct)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{
		"message": "Operação realizada com sucesso.",
		"code": 00,
		"authorization_id": %d,
		"balance": {
			"amount": %d,
			"currency_code": 986
		}
	}`, id, amount)))
}
