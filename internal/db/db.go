package db

import (
	"log"

	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := ""

	if dsn == "" {
		log.Fatal("dsn not set")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	//auto migrate models
	err = db.AutoMigrate(&models.Customer{}, &models.Branch{}, &models.Account{}, &models.Transaction{}, &models.LoanPayment{}, &models.Beneficiary{})
	if err != nil {
		log.Fatal("automigrated failed", err)
	}
	return db
}
