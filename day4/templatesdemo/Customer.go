package main

import (
	"html/template"
	"log"
	"net/http"
)

type Customer struct {
	CustomerName string
	BankName     string
	Address      string
}

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "7070"
)

//function
func renderTemplate(w http.ResponseWriter, r *http.Request) {
	customer := Customer{"Parameswari",
		"HSBC", "Chennai"}
	parsedTemplate, _ := template.ParseFiles("templates/index.html")
	err := parsedTemplate.Execute(w, customer)
	if err != nil {
		log.Printf("Error occurred while executing the templateor writing its output : ", err)
		return
	}
}
func main() {
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
