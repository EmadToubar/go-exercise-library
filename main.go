package main

import (
	"fmt"
	"log"
	"os"

	"librarydatabase/db"
	"librarydatabase/models"
	"librarydatabase/printer"
	"librarydatabase/repositories"
	"librarydatabase/services"

	_ "github.com/lib/pq"
)

func main() {
	ctx := db.InitDb()
	pr := repositories.CreatePersonRepo(ctx)
	ps := services.NewPersonService(pr)
	option := 0
	printer.Menu()           //Print out the menu
	fmt.Scanf("%v", &option) //User input for options
	for option != 12 {       //KH: Loop until user input 12 to quit the program
		switch option {
		case 1:
			//Registering new borrower
			ps.AddPerson()
		case 2:
			//Adding new book
			title := ""
			isbn := ""
			status := false
			fmt.Println("Please insert the book's title and ISBN: ")
			fmt.Scan(&title, &isbn)
			tx := ctx.MustBegin()
			tx.MustExec("INSERT INTO book (title, isbn, status) VALUES ($1, $2, $3)", title, isbn, status)
			tx.Commit()
			println("Book added.")

		case 3:
			//Deleting a borrower by ID
			id := 0
			fmt.Println("Please input the ID of the borrower ")
			fmt.Scan(&id)
			ctx.Query("DELETE FROM person WHERE personID=($1)", id)
			println("Person deleted.")

		case 4:
			//Deleting a book by ISBN
			isbn_check := ""
			fmt.Println("Please input the ISBN of the book: ")
			fmt.Scan(&isbn_check)
			ctx.Query("DELETE FROM person WHERE isbn=($1)", isbn_check)
			println("Person deleted.")

		case 5:
			//Listing all borrowers

			pop := models.Person{}
			rows, err := ctx.Queryx("SELECT * FROM person")
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
			pop := models.Book{}
			rows, err := ctx.Queryx("SELECT * FROM book")
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
			pop := models.Loan{}
			rows, err := ctx.Queryx("SELECT * FROM person_book")
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
			people := models.Person{}
			err = ctx.Get(&people, "SELECT * FROM person WHERE personID=($1)", id)

			fmt.Printf("%#v\n", people)

		case 9:
			//Searching for a book
			isbn_check := ""
			fmt.Println("Please input the ISBN of the book: ")
			fmt.Scan(&isbn_check)
			books := []models.Book{}
			ctx.Select(&books, "SELECT * FROM book WHERE isbn=($1)", isbn_check)
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
			tx := ctx.MustBegin()
			tx.MustExec("INSERT INTO person_book (isbn, last_name) VALUES ($1, $2)", isbn_check, lastname)
			tx.Commit()

			ctx.Query("UPDATE book SET status=true WHERE isbn=($1)", isbn_check)
			println("Book loaned.")

		case 11:
			//Returning a book
			isbn_check := ""
			fmt.Println("Please input the ISBN of the book: ")
			fmt.Scan(&isbn_check)
			ctx.Query("UPDATE book SET status=false WHERE isbn=($1)", isbn_check)
			println("Book returned.")
		case 12:
			//Quitting system
			os.Exit(1)
		}
	}

}
