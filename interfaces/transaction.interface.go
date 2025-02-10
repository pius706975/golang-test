package interfaces

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/golang-test/package/database/models"
)

type TransactionRepo interface {
	CreateTransaction(trData *models.Transaction) (*models.Transaction, error)
	GetTransactions(userId string, status string, page int, limit int) ([]models.Transaction, int64, error)
	GetByID(transactionId string) (*models.Transaction, error)
	UpdateTransaction(transactionId string, newStatus string) (*models.Transaction, error)
	DeleteTransaction(transactionId string) error
	GetTransactionSummaryToday() (map[string]interface{}, error)
}

type TransactionService interface {
	CreateTransaction(userId string, amount float64, paymentMethod string) (gin.H, int)
	GetTransactions(userId string, status string, page int, limit int) (gin.H, int)
	GetByID(transactionId string) (gin.H, int)
	UpdateTransaction(transactionId string, newStatus string) (gin.H, int)
	DeleteTransaction(transactionId string) (gin.H, int)
	GetTransactionSummaryToday() (gin.H, int)
}
