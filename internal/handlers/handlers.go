package handlers

import (
	"net/http"
	"strconv"

	"github.com/dmitryzhvinklis/bank-test-task/accounts"

	"github.com/gin-gonic/gin"
)

func CreateAccountHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	account := accounts.CreateAccount(id)
	c.JSON(http.StatusOK, account)
}

func DepositHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
	err := accounts.Deposit(id, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func WithdrawHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
	err := accounts.Withdraw(id, amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func GetBalanceHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	balance, err := accounts.GetBalance(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
