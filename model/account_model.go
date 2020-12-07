package model

type Account struct {
	AccountNumber string `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int    `json:"balance"`
}

type Transfer struct {
	MyAccountNumber string `json:"my_account_number"`
	ToAccountNumber string `json:"to_account_number"`
	Amount          int    `json:"amount"`
}
