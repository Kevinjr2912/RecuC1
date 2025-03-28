package application

import (
	"examc1/domain/entities"
	"examc1/domain/repositories"
)

type CreatePerson struct {
	db repositories.IPerson
}


func NewCreatePerson(db repositories.IPerson) *CreatePerson {
	return &CreatePerson{db: db}
}

func (cp * CreatePerson) Run(person entities.Person) error {
	return cp.db.CreatePerson(person)
}