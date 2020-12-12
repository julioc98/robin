package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/julioc98/robin/cmd/api/handler"
	"github.com/julioc98/robin/cmd/api/router"
	"github.com/julioc98/robin/internal/app/account"
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

	conn := db.Conn()
	defer conn.Close()
	db.Migrate(conn)

	r := mux.NewRouter()
	r.Use(middleware.Logging)

	accountRep := account.NewPostgresRepository(conn)
	accountService := account.NewService(accountRep)
	accountHandler := handler.NewAccountHandler(accountService)

	router.SetAccountRoutes(accountHandler, r.PathPrefix("/accounts").Subrouter())

	transactionRep := transaction.NewPostgresRepository(conn)
	transactionService := transaction.NewService(transactionRep)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	router.SetTransactionRoutes(transactionHandler, r.PathPrefix("/transactions").Subrouter())

	var integration struct{}

	integrationHandler := handler.NewIntegrationHandler(integration)
	router.SetIntegrationRoutes(integrationHandler, r.PathPrefix("/integrations").Subrouter())

	r.HandleFunc("/", handlerHi)
	http.Handle("/", r)

	port := env.Get("PORT", "5001")
	log.Printf(`%s listening on port: %s `, env.Get("APP", "robin"), port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
