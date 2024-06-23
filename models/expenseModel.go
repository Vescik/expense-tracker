package models

import (
	"encoding/json"
	"expense-tracker/config"
	"fmt"
)

// Properly reference the models package

type Expense struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	TypeID        int8    `json:"type_id"`
	Amount        float32 `json:"amount"`
	TransactionID int8    `json:"transaction_id"`
}
type ExpenseType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Transaction struct {
	Month     int  `json:"month"`
	Week      int8 `json:"week"`
	ExpenseID int  `json:"expense_id"`
}

type Rate struct {
	ID         int     `json:"id"`
	UpdatedAt  string  `json:"updated_at"`
	CurrencyID int     `json:"currency_id"`
	RateVal    float32 `json:"rate_val"`
}

func GetRates() {
	DbClient := config.CreateClient()
	data, count, err := DbClient.From("rates").Select("*", "exact", false).Execute()
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(json.Unmarshal(data))
	fmt.Print(count)
}
