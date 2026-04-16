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
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID   uint          `json:"customer_id" gorm:"type:int;index"`
	Customer     *Customer     `gorm:"foreignKey:CustomerID" json:"customer"`
	BranchID     uint          `json:"branch_id" gorm:"type:int;index"`
	Branch       *Branch       `gorm:"foreignKey:BranchID" json:"branch"`
	Amount       float64       `json:"amount" gorm:"type:float;not null"`
	InterestRate float64       `json:"interest_rate" gorm:"type:float;not null"`
	TermsMonths  uint          `json:"terms_months" gorm:"type:int;index"`
	TotalPayable float64       `json:"total_payable" gorm:"type:float;not null"`
	Status       string        `json:"status"`
	StartDate    time.Time     `json:"start_date"`
	EndDate      time.Time     `json:"end_date"`
	CreatedAt    time.Time     `json:"created_at"`
	Payments     []LoanPayment `gorm:"foreignKey:LoanID" json:"payments"`
}

type LoanPayment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	LoanID    uint      `json:"loan_id" gorm:"type:int;index"`
	Loan      *Loan     `gorm:"foreignKey:LoanID" json:"loan"`
	Amount    float64   `json:"amount" gorm:"type:decimal(15,2)"`
	DueDate   time.Time `json:"due_date" gorm:"type:date"`
	PaidDate  time.Time `json:"paid_date" gorm:"type:date"`
	Status    string    `json:"status" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"created_at" gorm:"type:date"`
}

type Beneficiary struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID    uint      `json:"customer_id" gorm:"type:int;index"`
	Customer      *Customer `gorm:"foreignKey:CustomerID" json:"customer"`
	Name          string    `json:"name" gorm:"type:varchar(100)"`
	AccountNumber string    `json:"account_number" gorm:"type:varchar(100)"`
	BankName      string    `json:"bank_name" gorm:"type:varchar(100)"`
}

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required"`
}

type TransferRequest struct {
	ToAccountID uint    `json:"to_account_id" binding:"required"`
	Amount      float64 `json:"amount" binding:"required, gt=0"`
}

type CreateLoanRequest struct {
	Amount       float64 `json:"amount" binding:"required, gt=0"`
	InterestRate float64 `json:"interest_rate" binding:"required, gt=0"`
	TermsMonths  uint    `json:"terms_months" binding:"required"`
}

type RegisterCustomerRequest struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required, min=8"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

type AddBeneficiaryRequest struct {
	Name          string `json:"name" binding:"required"`
	AccountNumber string `json:"account_number" binding:"required"`
	BankName      string `json:"bank_name"`
}

type DepositRequest struct {
	Amount float64 `json:"amount" binding:"required, gt=0"`
}

type MakePaymentRequest struct {
	PaymentID uint `json:"payment_id" binding:"required"`
}
