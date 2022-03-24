package repositories

import (
	"fmt"
	"librarydatabase/models"

	"github.com/jmoiron/sqlx"
)

type personRepo struct {
	db *sqlx.DB
}

type PersonRepo interface {
	AddPerson(models.Person)
	GetPersonByID(id int) (models.Person, error)
}

func CreatePersonRepo(db *sqlx.DB) PersonRepo {
	return &personRepo{db: db}
}

func (p *personRepo) GetPersonByID(id int) (models.Person, error) {
	var person models.Person
	err := p.db.Get(&person, "SELECT * FROM person WHERE personid=$1", id)
	return person, err
}

func (p *personRepo) AddPerson(person models.Person) {
	firstname := ""
	lastname := ""
	email := ""
	fmt.Println("Please insert your first name, last name, and email: ")
	fmt.Scan(&firstname, &lastname, &email)
	tx := p.db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", firstname, lastname, email)
	tx.Commit()
	println("User created.")
}
