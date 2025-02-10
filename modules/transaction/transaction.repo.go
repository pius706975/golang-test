package transaction

import (
	"encoding/json"
	"time"

	"github.com/pius706975/golang-test/package/database/models"
	"gorm.io/gorm"
)

type transactionRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *transactionRepo {
	return &transactionRepo{db}
}

func (repo *transactionRepo) CreateTransaction(trData *models.Transaction) (*models.Transaction, error) {
	statusHistory := []models.StatusHistory{
		{
			Status:    "pending",
			Timestamp: time.Now(),
		},
	}

	jsonData, _ := json.Marshal(statusHistory)
	trData.StatusHistory = jsonData
	trData.Status = "pending"

	if err := repo.db.Create(trData).Error; err != nil {
		return nil, err
	}

	return trData, nil
}

func (repo *transactionRepo) UpdateTransaction(transactionId string, newStatus string) (*models.Transaction, error) {
	var transaction models.Transaction

	err := repo.db.Where("id = ?", transactionId).First(&transaction).Error
	if err != nil {
		return nil, err
	}

	transaction.Status = newStatus

	var statusHistory []models.StatusHistory
	json.Unmarshal(transaction.StatusHistory, &statusHistory)
	statusHistory = append(statusHistory, models.StatusHistory{
		Status:    newStatus,
		Timestamp: time.Now(),
	})
	transaction.StatusHistory, _ = json.Marshal(statusHistory)

	if err := repo.db.Save(&transaction).Error; err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (repo *transactionRepo) DeleteTransaction(transactionId string) error {
	var trData models.Transaction

	if err := repo.db.Delete(trData, "id = ?", transactionId).Error; err != nil {
		return err
	}

	return nil
}

func (repo *transactionRepo) GetTransactions(userId string, status string, page, limit int) ([]models.Transaction, int64, error) {
	var transactions []models.Transaction
	var totalRecords int64

	query := repo.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, username, email, phone_number, is_verified")
	}).Order("created_at DESC")

	if userId != "" {
		query = query.Where("user_id = ?", userId)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Model(&models.Transaction{}).Count(&totalRecords).Error
	if err != nil {
		return nil, 0, err
	}

	if limit > 0 && page > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	if err := query.Find(&transactions).Error; err != nil {
		return nil, 0, err
	}

	return transactions, totalRecords, nil
}

func (repo *transactionRepo) GetByID(transactionId string) (*models.Transaction, error) {
	var transaction models.Transaction

	err := repo.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, username, email, phone_number, is_verified")
	}).Where("id = ?", transactionId).First(&transaction).Error

	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (repo *transactionRepo) GetTransactionSummaryToday() (map[string]interface{}, error) {
	var totalSuccessTransactions int64
	var totalFailedTransactions int64
	var totalTransactionAmount float64
	var uniqueUsersCount int64
	var mostUsedPaymentMethod string
	var latestTransactions []models.Transaction
	var totalTransactions int64
	var averageTransactionPerUser float64

	startOfDay := time.Now().Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	err := repo.db.Model(&models.Transaction{}).
		Where("status = ? AND created_at BETWEEN ? AND ?", "success", startOfDay, endOfDay).
		Count(&totalSuccessTransactions).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&models.Transaction{}).
		Where("status = ? AND created_at BETWEEN ? AND ?", "failed", startOfDay, endOfDay).
		Count(&totalFailedTransactions).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&models.Transaction{}).
		Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).
		Select("COALESCE(SUM(amount), 0)").
		Scan(&totalTransactionAmount).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&models.Transaction{}).
		Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).
		Select("COUNT(DISTINCT user_id)").
		Scan(&uniqueUsersCount).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&models.Transaction{}).
		Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).
		Select("payment_method").
		Group("payment_method").
		Order("COUNT(payment_method) DESC").
		Limit(1).
		Scan(&mostUsedPaymentMethod).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name, email, phone_number, created_at, updated_at") 
	}).
		Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).
		Order("created_at DESC").
		Limit(10).
		Find(&latestTransactions).Error
	if err != nil {
		return nil, err
	}

	err = repo.db.Model(&models.Transaction{}).
		Where("created_at BETWEEN ? AND ?", startOfDay, endOfDay).
		Count(&totalTransactions).Error
	if err != nil {
		return nil, err
	}

	if uniqueUsersCount > 0 {
		averageTransactionPerUser = float64(totalTransactions) / float64(uniqueUsersCount)
	} else {
		averageTransactionPerUser = 0
	}

	summary := map[string]interface{}{
		"total_success_transactions":    totalSuccessTransactions,
		"total_failed_transactions":     totalFailedTransactions,
		"total_transaction_amount":      totalTransactionAmount,
		"unique_users_count":            uniqueUsersCount,
		"most_used_payment_method":      mostUsedPaymentMethod,
		"latest_transactions":           latestTransactions,
		"average_transactions_per_user": averageTransactionPerUser,
	}

	return summary, nil
}
