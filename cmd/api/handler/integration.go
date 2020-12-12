package handler

import (
	"net/http"
)

// IntegrationHandler Interface
type IntegrationHandler interface {
	Status(w http.ResponseWriter, r *http.Request)
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

// Add a Integration
func (ah *integrationHandler) Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{
		"message": "Operação realizada com sucesso.",
		"code": 0
	  }`))
}
