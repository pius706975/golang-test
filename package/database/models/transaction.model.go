package models

import (
	"time"

	"gorm.io/datatypes"
)

type StatusHistory struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

type Transaction struct {
	ID string `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`

	UserID string `gorm:"not null" json:"user_id" valid:"uuidv4"`
	User   User   `gorm:"foreignKey:UserID" json:"user_data" valid:"-"`

	Amount        float64        `gorm:"not null" json:"amount,omitempty" valid:"type(float64), required~Amount is required"`
	Status        string         `gorm:"not null" json:"status,omitempty" valid:"type(string)"`
	StatusHistory datatypes.JSON `gorm:"type:jsonb" json:"status_history"`
	PaymentMethod string         `gorm:"not null" json:"payment_method,omitempty" valid:"type(string), required~Payment method is required"`
	CreatedAt     time.Time      `json:"created_at"  valid:"-"`
	UpdatedAt     time.Time      `json:"updated_at" valid:"-"`
}

type Transactions []Transaction

func (Transaction) TableName() string {
	return "transactions"
}

type TransactionResponse struct {
	ID            string         `json:"id"`
	UserID        string         `json:"user_id"`
	Amount        float64        `json:"amount"`
	Status        string         `json:"status"`
	StatusHistory datatypes.JSON `json:"status_history"`
	PaymentMethod string         `json:"payment_method"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type CreateTransactionRequest struct {
	Amount        float64 `json:"amount" valid:"type(float64), required~Amount is required"`
	PaymentMethod string  `json:"payment_method" valid:"type(string), required~Payment method is required"`
}

type UpdateTransactionStatusRequest struct {
	Status string `json:"status" valid:"type(string), required~Status is required"`
}
