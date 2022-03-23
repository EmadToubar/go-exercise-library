package main

import (
	"fmt"
	"log"
	"os"

	"librarydatabase/printer"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Person struct {
	PersonID  int    `db:"personid"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Borrowed  []Book `db:"borrowed"`
}

type Book struct {
	Title           string `db:"title"`
	ISBN            string `db:"isbn"`
	Borrowed_status bool   `db:"status"`
}

type Loan struct {
	ISBN     string `db:"isbn"`
	LastName string `db:"last_name"`
}

var schema = `

CREATE TABLE IF NOT EXISTS person (
    personID SERIAL PRIMARY KEY,
	first_name text,
    last_name text,
    email text
);


CREATE TABLE IF NOT EXISTS book (
	title text,
	isbn text PRIMARY KEY,
	status boolean
	);

CREATE TABLE IF NOT EXISTS person_book (
	loan_isbn text,
	loan_last_name text,
	FOREIGN KEY (loan_isbn) REFERENCES book(isbn) ON DELETE CASCADE,
	FOREIGN KEY (loan_last_name) REFERENCES person(last_name) ON DELETE CASCADE

			);
	
	`

func main() {

	db, err := sqlx.Connect("postgres", "user=postgres dbname=testdatabase password=emadsql sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	} //Connecting to database

	db.MustExec(schema)

	option := 0
	printer.Menu()           //Print out the menu
	fmt.Scanf("%v", &option) //User input for options

	switch option {
	case 1:
		//Registering new borrower
		firstname := ""
		lastname := ""
		email := ""
		fmt.Println("Please insert your first name, last name, and email: ")
		fmt.Scan(&firstname, &lastname, &email)
		tx := db.MustBegin()
		tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", firstname, lastname, email)
		tx.Commit()
		println("User created.")
	case 2:
		//Adding new book
		title := ""
		isbn := ""
		status := false
		fmt.Println("Please insert the book's title and ISBN: ")
		fmt.Scan(&title, &isbn)
		tx := db.MustBegin()
		tx.MustExec("INSERT INTO book (title, isbn, status) VALUES ($1, $2, $3)", title, isbn, status)
		tx.Commit()
		println("Book added.")

	case 3:
		//Deleting a borrower by ID
		id := 0
		fmt.Println("Please input the ID of the borrower ")
		fmt.Scan(&id)
		db.Query("DELETE FROM person WHERE personID=($1)", id)
		println("Person deleted.")

	case 4:
		//Deleting a book by ISBN
		isbn_check := ""
		fmt.Println("Please input the ISBN of the book: ")
		fmt.Scan(&isbn_check)
		db.Query("DELETE FROM person WHERE isbn=($1)", isbn_check)
		println("Person deleted.")

	case 5:
		//Listing all borrowers

		pop := Person{}
		rows, err := db.Queryx("SELECT * FROM person")
		if err == err {
		}
		for rows.Next() {
			err := rows.StructScan(&pop)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%#v\n", pop)
		}

	case 6:
		//Listing all books
		pop := Book{}
		rows, err := db.Queryx("SELECT * FROM book")
		if err == err {
		}
		for rows.Next() {
			err := rows.StructScan(&pop)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%#v\n", pop)
		}
	case 7:
		//Listing all loans
		pop := Loan{}
		rows, err := db.Queryx("SELECT * FROM person_book")
		if err == err {
		}
		for rows.Next() {
			err := rows.StructScan(&pop)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("%#v\n", pop)
		}
	case 8:
		//Searching for a borrower by ID
		//lastname := ""
		var id int
		fmt.Println("Please input the ID of the borrower: ")
		fmt.Scan(&id)
		people := Person{}
		err = db.Get(&people, "SELECT * FROM person WHERE personID=($1)", id)

		fmt.Printf("%#v\n", people)

	case 9:
		//Searching for a book
		isbn_check := ""
		fmt.Println("Please input the ISBN of the book: ")
		fmt.Scan(&isbn_check)
		books := []Book{}
		db.Select(&books, "SELECT * FROM book WHERE isbn=($1)", isbn_check)
		result := books[0]
		fmt.Printf("%#v\n", result)
	case 10:
		//Loaning a book
		isbn_check := ""
		lastname := ""
		fmt.Println("Please input the ISBN of the book: ")
		fmt.Scan(&isbn_check)
		fmt.Println("Please input the last name of the borrower ")
		fmt.Scan(&lastname)
		tx := db.MustBegin()
		tx.MustExec("INSERT INTO person_book (isbn, last_name) VALUES ($1, $2)", isbn_check, lastname)
		tx.Commit()

		db.Query("UPDATE book SET status=true WHERE isbn=($1)", isbn_check)
		println("Book loaned.")

	case 11:
		//Returning a book
		isbn_check := ""
		fmt.Println("Please input the ISBN of the book: ")
		fmt.Scan(&isbn_check)
		db.Query("UPDATE book SET status=false WHERE isbn=($1)", isbn_check)
		println("Book loaned.")
	case 12:
		//Quitting system
		os.Exit(1)
	}

}
