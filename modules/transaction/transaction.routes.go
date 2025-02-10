package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/pius706975/golang-test/middlewares"
)

func TransactionRoutes(router *gin.Engine, controller *transactionController, prefix string) {

	transactionGroup := router.Group(prefix + "/transactions")
	{
		transactionGroup.POST("", middlewares.AuthMiddleware(), func(ctx *gin.Context) {
			controller.CreateTransaction(ctx)
		})

		transactionGroup.GET("", func(ctx *gin.Context) {
			controller.GetTransactions(ctx)
		})

		transactionGroup.GET("/:id", func(ctx *gin.Context) {
			controller.GetByID(ctx)
		})

		transactionGroup.GET("/dashboard/summary", func(ctx *gin.Context) {
			controller.GetTransactionSummaryToday(ctx)
		})

		transactionGroup.PUT("/:id", func(ctx *gin.Context) {
			controller.UpdateTransaction(ctx)
		})

		transactionGroup.DELETE("/:id", func(ctx *gin.Context) {
			controller.DeleteTransaction(ctx)
		})
	}
}
