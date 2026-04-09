package models

import "time"

type Customer struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName     string    `gorm:"unique;type:varchar(255);not null" json:"username"`
	PassWordHash string    `gorm:"type:varchar(255);not null" json:"-"`
	FirstName    string    `gorm:"type:varchar(100)" json:"first_name"`
	LastName     string    `gorm:"type:varchar(100)" json:"last_name"`
	Email        string    `gorm:"unique;type:varchar(255);not null" json:"email"`
	Phone        string    `gorm:"unique;type:varchar(20)" json:"phone"`
	Address      string    `gorm:"type:text" json:"address"`
	CreatedAt    time.Time `json:"created_at"`

	Accounts      []Account     `gorm:"foreignKey:CustomerID" json:"-"`
	Loans         []Loan        `gorm:"foreignKey:CustomerID" json:"-"`
	Beneficiaries []Beneficiary `gorm:"foreignKey:CustomerID" json:"-"`
}

type Branch struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	Code    string `gorm:"unique;type:varchar(50);not null" json:"code"`
	City    string `gorm:"type:varchar(100)" json:"city"`
	Address string `gorm:"type:text" json:"address"`
	Phone   string `gorm:"type:varchar(20)" json:"phone"`

	Accounts []Account `gorm:"foreignKey:BranchID" json:"-"`
	Loans    []Loan    `gorm:"foreignKey:BranchID" json:"-"`
}

type Account struct {
	ID         uint     `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID uint     `json:"customer_id" gorm:"type:int;index"`
	Customer   Customer `gorm:"foreignKey:CustomerID" json:"customer"`

	BranchID uint   `json:"branch_id" gorm:"type:int;index"`
	Branch   Branch `gorm:"foreignKey:BranchID" json:"branch"`

	Owner    string  `json:"owner" gorm:"type:varchar(100)"`
	Balance  float64 `json:"balance" gorm:"type:float;not null"`
	Currency string  `json:"currency" gorm:"type:varchar(100)"`

	CreatedAt    time.Time     `json:"created_at"`
	Transactions []Transaction `gorm:"foreignKey:FromAccountID;references:ID" json:"-"`
}

type Transaction struct {
	ID            uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	FromAccountID uint    `json:"from_account_id" gorm:"type:int;index"`
	ToAccountID   uint    `json:"to_account_id" gorm:"type:int;index"`
	LoanPaymentID uint    `json:"loan_payment_id" gorm:"type:int;index"`
	BeneficiaryID uint    `json:"beneficiary_id" gorm:"type:int;index"`
	Amount        float64 `json:"amount" gorm:"type:float;not null"`

	CreatedAt time.Time `json:"created_at"`
}

type Loan struct {
	ID           uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID   uint      `json:"customer_id" gorm:"type:int;index"`
	Customer     *Customer `gorm:"foreignKey:CustomerID" json:"customer"`
	BranchID     uint      `json:"branch_id" gorm:"type:int;index"`
	Branch       *Branch   `gorm:"foreignKey:BranchID" json:"branch"`
	Amount       float64   `json:"amount" gorm:"type:float;not null"`
	InterestRate float64   `json:"interest_rate" gorm:"type:float;not null"`
}
