package routers

import (
	"github.com/dmitryzhvinklis/bank-test-task/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/accounts", handlers.CreateAccountHandler)
	r.POST("/accounts/:id/deposit", handlers.DepositHandler)
	r.POST("/accounts/:id/withdraw", handlers.WithdrawHandler)
	r.GET("/accounts/:id/balance", handlers.GetBalanceHandler)
	return r
}
