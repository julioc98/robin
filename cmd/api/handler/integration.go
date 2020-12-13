package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julioc98/robin/internal/app/entity"
)

// IntegrationHandler Interface
type IntegrationHandler interface {
	Status(w http.ResponseWriter, r *http.Request)
	AddPurchase(w http.ResponseWriter, r *http.Request)
}

// IntegrationService Interface
type IntegrationService interface{}

type integrationHandler struct {
	integrationService IntegrationService
}

// NewIntegrationHandler Create a new handler
func NewIntegrationHandler(integrationService IntegrationService) IntegrationHandler {
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

	// id, err := ah.accountService.Create(&req)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	authCode, amount := 322169, 200000

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
	}`, authCode, amount)))
}
