package models

type Loan struct {
	ISBN     string `db:"isbn"`
	LastName string `db:"last_name"`
}
