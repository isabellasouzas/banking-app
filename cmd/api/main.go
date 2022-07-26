package main

import (
	"context"
	"log"

	"banking-app/app"
	"banking-app/app/config"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadConfigurations()
	if err != nil {
		log.Fatalf("Could not load environment: %v", err)
	}

	application, err := app.NewApplication(ctx, cfg)
	if err != nil {
		log.Fatalf("Could not set up application: %v", err)
	}

	//router := mux.NewRouter()
	//
	//// '/accounts'
	//router.HandleFunc("/accounts", Handler.reate).Methods("GET")
	//router.HandleFunc("/accounts/{id}/", http.HandleGetAccountByID).Methods("GET")
	//router.HandleFunc("/accounts", Create).Methods("POST")
	//
	//// transfers
	//router.HandleFunc("/transfer/{account}", http.HandleGetTransfers).Methods("GET")
	//router.HandleFunc("/transfer/{account}/{destination}/{amount}", http.HandleSendTransfer).Methods("POST")
	//
	//fmt.Printf("Server is running at %s", port)
	//log.Fatal(http.ListenAndServe(port, router))
}
