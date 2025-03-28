package repositories

import "examc1/domain/entities"

type IPerson interface {
	CreatePerson(person entities.Person) error
	GetAllPersons() ([]entities.Person, error)
	CountGender() ([]int, error)
}