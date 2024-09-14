package loan_example

import (
	"errors"
)

type Loan interface {
	GetLoanId() string
	GetAmount() float64
	GetIncome() float64
	GetInstallments() int
	GetLoanType() string
	GetRate() float64
}

type MortgageLoan struct {
	loanId       string
	amount       float64
	income       float64
	installments int
	loanType     string
	rate         float64
}

func (m *MortgageLoan) GetLoanId() string    { return m.loanId }
func (m *MortgageLoan) GetAmount() float64   { return m.amount }
func (m *MortgageLoan) GetIncome() float64   { return m.income }
func (m *MortgageLoan) GetInstallments() int { return m.installments }
func (m *MortgageLoan) GetLoanType() string  { return m.loanType }
func (m *MortgageLoan) GetRate() float64     { return m.rate }

type CarLoan struct {
	loanId       string
	amount       float64
	income       float64
	installments int
	loanType     string
	rate         float64
}

func (c *CarLoan) GetLoanId() string    { return c.loanId }
func (c *CarLoan) GetAmount() float64   { return c.amount }
func (c *CarLoan) GetIncome() float64   { return c.income }
func (c *CarLoan) GetInstallments() int { return c.installments }
func (c *CarLoan) GetLoanType() string  { return c.loanType }
func (c *CarLoan) GetRate() float64     { return c.rate }

type LoanFactory interface {
	CreateLoan(id string, amount, income float64, installments int) (Loan, error)
}

type MortgageLoanFactory struct{}

func (f *MortgageLoanFactory) CreateLoan(id string, amount, income float64, installments int) (Loan, error) {
	if installments > 420 {
		return nil, errors.New("the maximum number of installments for mortgage loan is 420")
	}
	if income*0.25 < amount/float64(installments) {
		return nil, errors.New("the maximum number of installments could not exceed 25% of monthly income")
	}
	return &MortgageLoan{
		loanId:       id,
		amount:       amount,
		income:       income,
		installments: installments,
		loanType:     "Mortgage",
		rate:         10.0,
	}, nil
}

type CarLoanFactory struct{}

func (f *CarLoanFactory) CreateLoan(id string, amount, income float64, installments int) (Loan, error) {
	if installments > 60 {
		return nil, errors.New("the maximum number of installments for car loan is 60")
	}
	if income*0.3 < amount/float64(installments) {
		return nil, errors.New("the maximum number of installments could not exceed 30% of monthly income")
	}
	return &CarLoan{
		loanId:       id,
		amount:       amount,
		income:       income,
		installments: installments,
		loanType:     "Car",
		rate:         10.0,
	}, nil
}
