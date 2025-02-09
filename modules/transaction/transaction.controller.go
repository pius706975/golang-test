package transaction

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/pius706975/golang-test/interfaces"
	"github.com/pius706975/golang-test/package/database/models"
)

type transactionController struct {
	service interfaces.TransactionService
}

func NewController(service interfaces.TransactionService) *transactionController {
	return &transactionController{service}
}

// CreateTransaction godoc
// @Summary Create Transaction
// @Description Create transaction
// @Tags Transactions
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param userData body models.CreateTransactionRequest true "Transaction data"
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /api/transactions [post]
func (controller *transactionController) CreateTransaction(ctx *gin.Context) {
	ctx.Header("content-type", "application/json")
	userId, exists := ctx.Get("id")

	var trData models.Transaction

	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	}

	err := ctx.ShouldBindJSON(&trData)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
		return
	}

	_, err = govalidator.ValidateStruct(&trData)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	responseData, status := controller.service.CreateTransaction(userId.(string), trData.Amount, trData.PaymentMethod)

	ctx.JSON(status, responseData)
}

// UpdateTransaction godoc
// @Summary Update Transaction
// @Description Update transaction status
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Param userData body models.UpdateTransactionStatusRequest true "Transaction data"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /api/transactions/{id} [put]
func (controller *transactionController) UpdateTransaction(ctx *gin.Context) {
	transactionId := ctx.Param("id")

	var requestBody struct {
		Status string `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(400, gin.H{"message": "Invalid request body"})
		return
	}

	responseData, statusCode := controller.service.UpdateTransaction(transactionId, requestBody.Status)
	ctx.JSON(statusCode, responseData)
}

// DeleteRole godoc
// @Summary Delete a transaction
// @Description Delete a transaction by role ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/transactions/{id} [delete]
func (controller *transactionController) DeleteTransaction(ctx *gin.Context) {
	id := ctx.Param("id")

	responseData, status := controller.service.DeleteTransaction(id)

	ctx.JSON(status, responseData)
}

// GetTransactions godoc
// @Summary Get Transactions
// @Description Retrieve transactions with optional filters (user_id, status, pagination)
// @Tags Transactions
// @Accept json
// @Produce json
// @Param user_id query string false "User ID filter"
// @Param status query string false "Transaction status filter (pending, success, failed)"
// @Param page query int false "Page number for pagination (default: 1)"
// @Param limit query int false "Limit number of transactions per page (default: 10)"
// @Success 200 {object} map[string]interface{} "Successful response"
// @Failure 400 {object} map[string]interface{} "Bad Request"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /api/transactions [get]
func (controller *transactionController) GetTransactions(ctx *gin.Context) {
	userID := ctx.Query("user_id")
	status := ctx.Query("status")

	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))   
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	responseData, statusCode := controller.service.GetTransactions(userID, status, page, limit)
	ctx.JSON(statusCode, responseData)
}


// GetById godoc
// @Summary Get transaction by ID
// @Description Fetch the transaction details based on the ID provided
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200
// @Failure 404
// @Failure 500
// @Router /api/transactions/{id} [get]
func (controller *transactionController) GetByID(ctx *gin.Context) {
	transactionId := ctx.Param("id")

	responseData, statusCode := controller.service.GetByID(transactionId)

	ctx.JSON(statusCode, responseData)
}

// GetTransactionSummary godoc
// @Summary Get Transactions summary
// @Description Retrieve transactions summary
// @Tags Transactions
// @Accept json
// @Produce json
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/transactions/dashboard/summary [get]
func (controller *transactionController) GetTransactionSummaryToday(ctx *gin.Context) {
	responseData, statusCode := controller.service.GetTransactionSummaryToday()
	
	ctx.JSON(statusCode, responseData)
}
