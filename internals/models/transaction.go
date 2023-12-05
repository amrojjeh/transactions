package models

import (
	"database/sql"
	"time"
)

type Transaction struct {
	ID          int
	TDate       time.Time
	PDate       time.Time
	FromAccount string
	ToAccount   string
	Amount      float64
	Context     string
	Tags        string
}

type TransactionModel struct {
	DB *sql.DB
}

func (m TransactionModel) NewTransaction(tdate, pdate time.Time, from_account,
	to_account string, amount float64, context, tags string) error {
	stmt := `INSERT INTO transactions (tdate, pdate, from_account, to_account,
		amount, context, tags) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := m.DB.Exec(stmt, tdate, pdate, from_account, to_account, amount, context, tags)
	if err != nil {
		return err
	}
	return nil
}

func (m TransactionModel) GetTransactions() ([]Transaction, error) {
	stmt := `SELECT pdate, from_account, to_account, amount, tags FROM transactions`

	res, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	ts := []Transaction{}
	for res.Next() {
		t := Transaction{}
		err = res.Scan(&t.PDate, &t.FromAccount, &t.ToAccount, &t.Amount, &t.Tags)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}
