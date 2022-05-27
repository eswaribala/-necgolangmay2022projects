package models

import "time"

type Trader struct {
	TraderId   int64  `json:"traderId" gorm:"primary_key"`
	TraderName string `json:"traderName"`
	DOJ        time.Time
}
