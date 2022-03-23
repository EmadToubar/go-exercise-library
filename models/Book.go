package models

type Book struct {
	Title           string `db:"title"`
	ISBN            string `db:"isbn"`
	Borrowed_status bool   `db:"status"`
}
