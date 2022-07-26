package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	swaggerFiles "github.com/swaggo/files"
)

const port = ":8080"

// @title           Banking API
// @version         1.0
// @description     This is a Banking application, the solution to the management banking account.
//
// @contact.name   API Support
// @contact.email  isabsantos92@gmail.com

func main() {

	router := mux.NewRouter()

	// '/accounts'
	router.HandleFunc("/accounts", http.HandleAccounts).Methods("GET")
	router.HandleFunc("/accounts/{id}/", http.HandleGetAccountByID).Methods("GET")
	router.HandleFunc("/accounts", http.HandleCreateAccount).Methods("POST")

	// transfers
	router.HandleFunc("/transfer/{account}", http.HandleGetTransfers).Methods("GET")
	router.HandleFunc("/transfer/{account}/{destination}/{amount}", http.HandleSendTransfer).Methods("POST")

	router.Get("/swagger/*any", mux.ginSwagger.WrapHandler(swaggerFiles.Handler))

	fmt.Printf("Server is running at %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
