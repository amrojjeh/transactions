package models

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID           int
	Tdate        time.Time
	Pdate        time.Time
	from_account string
	to_account   string
	amount       float64
	context      string
	tags         []string
}

type TransactionModel struct {
	DB *sql.DB
}
