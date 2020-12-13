package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julioc98/robin/cmd/api/handler"
	"github.com/julioc98/robin/cmd/api/router"
	"github.com/julioc98/robin/internal/app/account"
	"github.com/julioc98/robin/internal/app/processor"
	"github.com/julioc98/robin/internal/app/transaction"
	db "github.com/julioc98/robin/internal/db"
	"github.com/julioc98/robin/pkg/env"
	"github.com/julioc98/robin/pkg/middleware"
)

func handlerHi(w http.ResponseWriter, r *http.Request) {
	msg := "Ola, Seja bem vindo ao Robin!!"
	log.Println(msg)
	w.Write([]byte(msg))
}

func main() {
	apiKey := env.Get("API_KEY", "")
	if apiKey == "" {
		log.Println("required API_KEY")
		return
	}

	conn := db.Conn()
	defer conn.Close()
	db.Migrate(conn)

	r := mux.NewRouter()
	r.Use(middleware.Logging)
	client := http.DefaultClient

	processorClient := processor.NewClient(apiKey, client)
	processorGatewayBasePath := env.Get("PAY_PATH", "https://api-hml.paysmart.com.br/paySmart/ps-processadora/v1")
	processorGateway := processor.NewPayGateway(processorGatewayBasePath, processorClient)

	accountRep := account.NewPostgresRepository(conn)
	accountService := account.NewService(accountRep, processorGateway)
	accountHandler := handler.NewAccountHandler(accountService)

	router.SetAccountRoutes(accountHandler, r.PathPrefix("/accounts").Subrouter())

	transactionRep := transaction.NewPostgresRepository(conn)
	transactionService := transaction.NewService(transactionRep)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	router.SetTransactionRoutes(transactionHandler, r.PathPrefix("/transactions").Subrouter())

	integrationHandler := handler.NewIntegrationHandler(transactionService)
	router.SetIntegrationRoutes(integrationHandler, r.PathPrefix("/integrations").Subrouter())

	r.HandleFunc("/", handlerHi)
	http.Handle("/", r)

	port := env.Get("PORT", "5001")
	log.Printf(`%s listening on port: %s `, env.Get("APP", "robin"), port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
