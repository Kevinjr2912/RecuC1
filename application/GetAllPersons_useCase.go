package application

import (
	"examc1/domain/entities"
	"examc1/domain/repositories"
)

type GetAllPerson struct {
	db repositories.IPerson
}

func NewGetAllPersons(db repositories.IPerson) *GetAllPerson {
	return &GetAllPerson{db: db}
}

func (gap *GetAllPerson) Run() ([]entities.Person, error) {
	return gap.db.GetAllPersons()
}