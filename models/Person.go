package models

type Person struct {
	PersonID  int    `db:"personid"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Borrowed  []Book `db:"borrowed"`
}
