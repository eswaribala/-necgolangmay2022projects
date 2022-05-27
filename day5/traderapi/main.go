package main

import (
	"github.com/gorilla/mux"
	_ "github.com/necws/day5/traderapi/docs"
	"github.com/necws/day5/traderapi/stores"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Traders API
// @version 1.0
// @description This is a sample service for managing Traders
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email param@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:7070
// @BasePath /
func main() {

	router := mux.NewRouter()
	// Create
	router.HandleFunc("/traders", stores.CreateTrader).Methods("POST")
	// Read
	router.HandleFunc("/traders/{traderId}", stores.GetTrader).Methods("GET")
	// Read-all
	router.HandleFunc("/traders", stores.GetTraders).Methods("GET")
	// Update
	router.HandleFunc("/traders/{traderId}", stores.UpdateTrader).Methods("PUT")
	// Delete
	router.HandleFunc("/traders/{traderId}", stores.DeleteTrader).Methods("DELETE")
	// Initialize db connection
	stores.InitDB()

	//log.Fatal(http.ListenAndServe(":7070", router))

	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Fatal(http.ListenAndServe(":7070", router))
}
