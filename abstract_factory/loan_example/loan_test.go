package loan_example

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoanFactory_ShouldCreateAMortgageLoan(t *testing.T) {
	loan, err := NewMortgageLoan("abcd-1234", 100000, 10000, 240)
	assert.Equal(t, loan.LoanDetails.loanId, "abcd-1234")
	assert.Equal(t, loan.LoanDetails.amount, float64(100000))
	assert.Equal(t, loan.LoanDetails.income, float64(10000))
	assert.Equal(t, loan.LoanDetails.installments, 240)
	assert.Nil(t, err)
}

func TestLoanFactory_ShouldNotCreateAMortgageLoanWithInstallmentsGreaterThan420(t *testing.T) {
	loan, err := NewMortgageLoan("abcd-1234", 100000, 10000, 450)
	assert.EqualError(t, err, "the maximum number of installments for mortgage loan is 420")
	assert.Nil(t, loan)
}

func TestLoanFactory_ShouldNotCreateAMortgageIfTheInstallmentOccupiesMoreThan25PercentOfTheMonthlyIncome(t *testing.T) {
	loan, err := NewMortgageLoan("abcd-1234", 200000, 1000, 240)
	assert.EqualError(t, err, "the maximum number of installments could not exceed 25% of monthly income")
	assert.Nil(t, loan)
}
