package transaction

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pius706975/golang-test/interfaces"
	"github.com/pius706975/golang-test/package/database/models"
)

type transactionService struct {
	repo interfaces.TransactionRepo
}

func NewService(repo interfaces.TransactionRepo) *transactionService {
	return &transactionService{repo}
}

func (service *transactionService) CreateTransaction(userID string, amount float64, paymentMethod string) (gin.H, int) {
	if amount < 10000 {
		return gin.H{"status": 400, "message": "Minimum transaction is 10,000"}, 400
	}

	transaction := &models.Transaction{
		UserID:        userID,
		Amount:        amount,
		PaymentMethod: paymentMethod,
		CreatedAt:     time.Now(),
	}

	tr, err := service.repo.CreateTransaction(transaction)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	responseData := models.TransactionResponse{
		ID:            tr.ID,
		UserID:        tr.UserID,
		Amount:        tr.Amount,
		Status:        tr.Status,
		StatusHistory: tr.StatusHistory,
		PaymentMethod: tr.PaymentMethod,
		CreatedAt:     tr.CreatedAt,
		UpdatedAt:     tr.UpdatedAt,
	}

	return gin.H{"status": 201, "message": "Transaction created successfully", "data": responseData}, 201
}

func (service *transactionService) UpdateTransaction(transactionId string, newStatus string) (gin.H, int) {
	allowedStatuses := map[string]bool{
		"pending": true,
		"success": true,
		"failed":  true,
	}

	if !allowedStatuses[newStatus] {
		return gin.H{"status": 400, "message": "Invalid transaction status"}, 400
	}

	transaction, err := service.repo.UpdateTransaction(transactionId, newStatus)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "message": "Transaction status updated", "data": transaction}, 200
}

func (service *transactionService) DeleteTransaction(transactionId string) (gin.H, int) {
	_, err := service.repo.GetByID(transactionId)
	if err != nil {
		return gin.H{"status": 404, "message": "Transaction not found"}, 404
	}

	err = service.repo.DeleteTransaction(transactionId)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "message": "Transaction deleted successfully"}, 200
}

func (service *transactionService) GetTransactions(userID, status string, page, limit int) (gin.H, int) {
	transactions, totalRecords, err := service.repo.GetTransactions(userID, status, page, limit)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{
		"status": 200,
		"data": gin.H{
			"transactions": transactions,
			"pagination": gin.H{
				"total_records": totalRecords,
				"current_page":  page,
				"limit":         limit,
				"total_pages":   int((totalRecords + int64(limit) - 1) / int64(limit)), 
			},
		},
	}, 200
}


func (service *transactionService) GetByID(transactionId string) (gin.H, int) {
	transaction, err := service.repo.GetByID(transactionId)
	if err != nil {
		return gin.H{"status": 404, "message": "Transaction not found"}, 404
	}

	return gin.H{"status": 200, "message": "Transaction detail fetched successfully", "data": transaction}, 200
}

func (service *transactionService) GetTransactionSummaryToday() (gin.H, int) {
	summary, err := service.repo.GetTransactionSummaryToday()
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "message": "Transaction summary retrieved", "data": summary}, 200
}
