package handlers

import (
	"bankapp/database"
	models "bankapp/pkg/model"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TxRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

var balances = make(map[uint]float64)

func Deposit(c *gin.Context) {
	var req TxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint)
	acctIDStr := c.Param("id")
	acctID, err := strconv.ParseUint(acctIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid account id"})
		return
	}

	log.Printf("[Deposit] userID=%d acctID=%d amount=%.2f", userID, acctID, req.Amount)

	res := database.DB.Debug(). // logs SQL
					Model(&models.Account{}).
					Where("id = ? AND user_id = ?", acctID, userID).
					UpdateColumn("balance", gorm.Expr("balance + ?", req.Amount))

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
		return
	}
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		return
	}

	var acct models.Account
	database.DB.First(&acct, acctID)
	c.JSON(http.StatusOK, acct)
}

func Withdraw(c *gin.Context) {
	var req TxRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint)
	acctID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		var acct models.Account
		if err := tx.First(&acct, "id = ? AND user_id = ?", acctID, userID).Error; err != nil {
			return err
		}
		if acct.Balance < req.Amount {
			return fmt.Errorf("insufficient funds")
		}
		return tx.Model(&acct).
			UpdateColumn("balance", gorm.Expr("balance - ?", req.Amount)).Error
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		} else if err.Error() == "insufficient funds" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	var acct models.Account
	database.DB.First(&acct, acctID)
	c.JSON(http.StatusOK, acct)
}
func GetBalance(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	var acct models.Account
	result := database.DB.
		First(&acct, "user_id = ?", userID)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id": userID,
		"balance": acct.Balance,
	})
}
