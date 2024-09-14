package loan_example

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoanFactory_ShouldCreateAMortgageLoan(t *testing.T) {
	loanFactory := &MortgageLoanFactory{}
	loan, err := loanFactory.CreateLoan("abcd-1234", 100000, 10000, 240)
	assert.Nil(t, err)
	assert.Equal(t, loan.GetLoanId(), "abcd-1234")
	assert.Equal(t, loan.GetAmount(), float64(100000))
	assert.Equal(t, loan.GetIncome(), float64(10000))
	assert.Equal(t, loan.GetInstallments(), 240)
}

func TestLoanFactory_ShouldNotCreateAMortgageLoanWithInstallmentsGreaterThan420Months(t *testing.T) {
	loanFactory := &MortgageLoanFactory{}
	loan, err := loanFactory.CreateLoan("abcd-1234", 100000, 10000, 450)
	assert.EqualError(t, err, "the maximum number of installments for mortgage loan is 420")
	assert.Nil(t, loan)
}

func TestLoanFactory_ShouldNotCreateAMortgageIfTheInstallmentOccupiesMoreThan25PercentOfTheMonthlyIncome(t *testing.T) {
	loanFactory := &MortgageLoanFactory{}
	loan, err := loanFactory.CreateLoan("abcd-1234", 200000, 1000, 240)
	assert.EqualError(t, err, "the maximum number of installments could not exceed 25% of monthly income")
	assert.Nil(t, loan)
}

func TestLoanFactory_ShouldNotCreateACarLoanWithInstallmentsGreaterThan60Months(t *testing.T) {
	loanFactory := &CarLoanFactory{}
	loan, err := loanFactory.CreateLoan("abcd-1234", 100000, 10000, 61)
	assert.EqualError(t, err, "the maximum number of installments for car loan is 60")
	assert.Nil(t, loan)
}

func TestLoanFactory_ShouldNotCreateACarIfTheInstallmentOccupiesMoreThan30PercentOfTheMonthlyIncome(t *testing.T) {
	loanFactory := &CarLoanFactory{}
	loan, err := loanFactory.CreateLoan("abcd-1234", 20000, 1000, 12)
	assert.EqualError(t, err, "the maximum number of installments could not exceed 30% of monthly income")
	assert.Nil(t, loan)
}
