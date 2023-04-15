package models

type Transaction struct {
	ID        int    `json:"id"`
	StartDate int    `json:"start_date"`
	DueDate   int    `json:"due_date"`
	UserID    int    `json:"user_id"`
	User      User   `json:"user"`
	Status    string `json:"status" gorm:"type: VARCHAR(25)"`
}

type TransactionResponse struct {
	ID        int    `json:"id"`
	StartDate int    `json:"start_date"`
	DueDate   int    `json:"due_date"`
	User      User   `json:"user"`
	Status    string `json:"status"`
}

func (TransactionResponse) TableName() string {
	return "TransactionResponse"
}
