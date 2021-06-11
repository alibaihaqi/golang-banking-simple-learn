package dto

type NewTransactionResponse struct {
	TransactionId string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}
