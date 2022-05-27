package stores

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/necws/day5/traderapi/models"
	"net/http"
	"strconv"
)

var db *gorm.DB

func InitDB() {
	var err error
	dataSourceName := "root:vignesh@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	//db.Exec("CREATE DATABASE nec_trader_db")
	db.Exec("USE nec_trader_db")

	// Migration to create tables for Order and Item schema
	db.AutoMigrate(&models.Trader{})
}

func CreateTrader(w http.ResponseWriter, r *http.Request) {
	var trader models.Trader
	json.NewDecoder(r.Body).Decode(&trader)
	// Creates new order by inserting records in the `orders` and `items` table
	db.Create(&trader)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trader)
}

func GetTrader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputTraderID := params["traderId"]

	var trader models.Trader
	db.First(&trader, inputTraderID)
	json.NewEncoder(w).Encode(trader)
}

func GetTraders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var traders []models.Trader
	db.Find(&traders)
	json.NewEncoder(w).Encode(traders)
}

func UpdateTrader(w http.ResponseWriter, r *http.Request) {
	var updatedTrader models.Trader
	json.NewDecoder(r.Body).Decode(&updatedTrader)
	db.Save(&updatedTrader)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTrader)
}

func DeleteTrader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputTraderID := params["traderId"]
	// Convert `orderId` string param to uint64
	id64, _ := strconv.ParseUint(inputTraderID, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("traderId = ?", idToDelete).Delete(&models.Trader{})
	w.WriteHeader(http.StatusNoContent)
}
