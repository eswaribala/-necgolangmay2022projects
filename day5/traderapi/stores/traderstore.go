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

	// Migration to create tables for Trader and Item schema
	db.AutoMigrate(&models.Trader{})
}

// CreateTrader godoc
// @Summary Create a new trader
// @Description Create a new trader with the input paylod
// @Tags traders
// @Accept  json
// @Produce  json
// @Param trader body models.Trader true "Create trader"
// @Success 200 {object} models.Trader
// @Router /traders [post]
func CreateTrader(w http.ResponseWriter, r *http.Request) {
	var trader models.Trader
	json.NewDecoder(r.Body).Decode(&trader)
	// Creates new trader by inserting records in the `traders` and `items` table
	db.Create(&trader)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(trader)
}

// GetTrader godoc
// @Summary Get details for a given traderId
// @Description Get details of order corresponding to the input traderId
// @Tags Traders
// @Accept  json
// @Produce  json
// @Param traderId path int true "ID of the trader"
// @Success 200 {object} models.Trader
// @Router /traders/{traderId} [get]

func GetTrader(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputTraderID := params["traderId"]

	var trader models.Trader
	db.First(&trader, inputTraderID)
	json.NewEncoder(w).Encode(trader)
}

// GetTrader godoc
// @Summary Get details of all traders
// @Description Get details of all traders
// @Tags traders
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Trader
// @Router /traders [get]

func GetTraders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var traders []models.Trader
	db.Find(&traders)
	json.NewEncoder(w).Encode(traders)
}

// UpdateTrader godoc
// @Summary Update trader identified by the given traderId
// @Description Update the trader corresponding to the input traderId
// @Tags traders
// @Accept  json
// @Produce  json
// @Param traderId path int true "ID of the trader to be updated"
// @Success 200 {object} models.Trader
// @Router /traders/{traderId} [post]

func UpdateTrader(w http.ResponseWriter, r *http.Request) {
	var updatedTrader models.Trader
	json.NewDecoder(r.Body).Decode(&updatedTrader)
	db.Save(&updatedTrader)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTrader)
}

// DeleteTrader godoc
// @Summary Delete trader identified by the given traderId
// @Description Delete the trader corresponding to the input traderId
// @Tags traders
// @Accept  json
// @Produce  json
// @Param traderId path int true "ID of the trader to be deleted"
// @Success 204 "No Content"
// @Router /traders/{traderId} [delete]

func DeleteTrader(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	inputTraderID := params["traderId"]
	// Convert `traderId` string param to uint64
	id64, _ := strconv.ParseUint(inputTraderID, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)

	db.Where("traderId = ?", idToDelete).Delete(&models.Trader{})
	w.WriteHeader(http.StatusNoContent)
}
