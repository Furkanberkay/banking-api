package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("cannot get current dir", err)
	}

	envPath := filepath.Join(dir, ".env")
	log.Printf("loading env vars from %s", envPath)

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal("loading .env file", err)
	} else {
		log.Println(".env file loaded")
	}

	dsn := os.Getenv("DB_DSN")
	jwtSecret := os.Getenv("JWT_SECRET")
	port := os.Getenv("PORT")

	if dsn == "" {
		log.Fatal("dsn not set")
	}
	if port == "" {
		port = "8080"
	}
	if jwtSecret == "" {
		log.Fatal("jwt secret not set")
	}

	_ = db.Connect()
	r := gin.Default()

	//cors middleware - applied globally for all routes
	r.Use(corsMiddleware())

	// public routes
	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)

	// protected routes
	protected := r.Group("")
	protected.Use(authMiddleWare())

	// accounts
	protected.POST("/accounts", handlers.CreateAccount)
	protected.GET("/accounts", handlers.ListAccounts)
	protected.POST("/transfers/:from_id", handlers.Transfer)
	protected.POST("/deposits/:account_id", handlers.Deposit)
	protected.GET("/accounts/:id/statements", handlers.GetStatements)

	// Loans
	protected.POST("/loans", handlers.CreateLoan)
	protected.GET("/loans", handlers.ListLoan)
	protected.POST("/loans/:id/repay", handlers.MakePayment)
	protected.POST("/loans/:id/payments", handlers.ListPayments)

	// beneficiaries
	protected.POST("/beneficiaries", handlers.AddBeneficiary)

	log.Printf("Server starting on %s", port)
	r.Run(":" + port)
}

func authMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("access-control-allow-Methods", "GET, POST, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set("access-control-allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("access-control-max-Age", "86400")
		c.Writer.Header().Set("access-control-allow-Credentials", "true")

		//handle preflight option request
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()

	}
}
