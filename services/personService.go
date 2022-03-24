package services

import (
	"fmt"
	"librarydatabase/models"
	"librarydatabase/repositories"
)

type personService struct {
	pr repositories.PersonRepo
}

type PersonService interface {
	AddPerson()
}

func NewPersonService(pr repositories.PersonRepo) PersonService {
	return &personService{pr: pr}
}

func (p *personService) AddPerson() {
	newPersonModel := models.Person{}
	fmt.Println("Please insert your first name, last name, and email: ")
	fmt.Scan(newPersonModel.FirstName, newPersonModel.LastName, newPersonModel.Email)
	p.pr.AddPerson(newPersonModel)
}
