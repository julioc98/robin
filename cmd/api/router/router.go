package router

import (
	"github.com/gorilla/mux"
	"github.com/julioc98/robin/cmd/api/handler"
)

// SetAccountRoutes add routes from Account
func SetAccountRoutes(ah handler.AccountHandler, r *mux.Router) {
	r.HandleFunc("", ah.Add).Methods("POST")
	r.HandleFunc("/auth", ah.FindByEmailAndPassword).Methods("POST")
	r.HandleFunc("/{id}", ah.FindByID).Methods("GET")
}

// SetTransactionRoutes add routes from Transaction
func SetTransactionRoutes(th handler.TransactionHandler, r *mux.Router) {
	r.HandleFunc("", th.Add).Methods("POST")
	r.HandleFunc("/{id:[0-9]+}", th.FindByID).Methods("GET")
	r.HandleFunc("/founds/{id}", th.GetFounds).Methods("GET")
}

// SetIntegrationRoutes add routes from Integrations
func SetIntegrationRoutes(th handler.IntegrationHandler, r *mux.Router) {
	r.HandleFunc("/status", th.Status).Methods("GET")
	r.HandleFunc("/purchases", th.AddPurchase).Methods("POST")
}
