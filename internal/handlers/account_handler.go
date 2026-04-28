package handlers

import (
	"banking-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandlers struct {
	accountSrv services.AccountService
}

func NewAccountHandler(s services.AccountService) *AccountHandlers {
	return &AccountHandlers{s}
}

func (ah *AccountHandlers) Register(c *gin.Context) {
	var req models.RegisterCustomerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ah.accountSrv.RegisterCustomer(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account created",
		"email":   req.Email,
	})
}

func Login(c *gin.Context) {
	var LoginReq struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&LoginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}
