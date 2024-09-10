package loan_example

import "errors"

type Loan interface {
	Rate() float64
}

type LoanDetails struct {
	loanId       string
	amount       float64
	income       float64
	installments int
	loanType     string
}

type MortgageLoan struct {
	*LoanDetails
	rate float64
}

func NewMortgageLoan(id string, amount, income float64, installments int) (*MortgageLoan, error) {
	if installments > 420 {
		return nil, errors.New("the maximum number of installments for mortgage loan is 420")
	}
	if income*0.25 < amount/float64(installments) {
		return nil, errors.New("the maximum number of installments could not exceed 25% of monthly income")
	}
	return &MortgageLoan{
		&LoanDetails{id, amount, income, installments, "Mortgage"},
		10,
	}, nil
}

func (m *MortgageLoan) Rate() float64 {
	return m.rate
}
