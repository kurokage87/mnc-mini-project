package model

import "github.com/google/uuid"

type RequestTransaction struct {
	TargetUser      uuid.UUID `json:"target_user"`
	Amount          float64   `json:"amount"`
	Remarks         string    `json:"remarks"`
	TransactionType string    `json:"transactionType"`
}

type DataResponseTopUp struct {
	TopUpId       uuid.UUID `json:"top_up_id"`
	AmountTopUp   float64   `json:"amount_top_up"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedDate   string    `json:"created_date"`
}

type DataResponsePayment struct {
	PaymentId     uuid.UUID `json:"payment_id"`
	Amount        float64   `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedDate   string    `json:"created_date"`
}

type DataResponseTransfer struct {
	TransferId    uuid.UUID `json:"transfer_id"`
	Amount        float64   `json:"amount"`
	Remarks       string    `json:"remarks"`
	BalanceBefore float64   `json:"balance_before"`
	BalanceAfter  float64   `json:"balance_after"`
	CreatedDate   string    `json:"created_date"`
}

type PaymentListTransactions struct {
	PaymentId       uuid.UUID `json:"payment_id"`
	Status          string    `json:"status"`
	UserId          uuid.UUID `json:"user_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceBefore   float64   `json:"balance_before"`
	BalanceAfter    float64   `json:"balance_after"`
	CreatedDate     string    `json:"created_date"`
}

type TopUpListTransactions struct {
	TopUpId         uuid.UUID `json:"top_up_id"`
	Status          string    `json:"status"`
	UserId          uuid.UUID `json:"user_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceBefore   float64   `json:"balance_before"`
	BalanceAfter    float64   `json:"balance_after"`
	CreatedDate     string    `json:"created_date"`
}

type TransferListTransactions struct {
	TransferId      uuid.UUID `json:"transfer_id"`
	Status          string    `json:"status"`
	UserId          uuid.UUID `json:"user_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          float64   `json:"amount"`
	Remarks         string    `json:"remarks"`
	BalanceBefore   float64   `json:"balance_before"`
	BalanceAfter    float64   `json:"balance_after"`
	CreatedDate     string    `json:"created_date"`
}

type ListTransactions struct {
	TransactionID       uuid.UUID `json:"transaction_id"`
	UserId              uuid.UUID `json:"user_id"`
	TransactionType     string    `json:"transaction_type"`
	TransactionCategory string    `json:"transaction_category"`
	Status              string    `json:"status"`
	Amount              float64   `json:"amount"`
	Remarks             string    `json:"remarks"`
	BalanceBefore       float64   `json:"balance_before"`
	BalanceAfter        float64   `json:"balance_after"`
	CreatedDate         string    `json:"created_date"`
}
