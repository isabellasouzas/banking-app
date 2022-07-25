package main

import (
	"fmt"
	"log"
	"net/http"

	_ "swag-gin-demo/docs"

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

	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}
	//defer conn.Close(context.Background())
	//
	//var greeting string
	//err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(greeting)

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
