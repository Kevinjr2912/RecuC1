package application

import "examc1/domain/repositories"

type GetCountGender struct {
	db repositories.IPerson
}

func NewGetCountGender(db repositories.IPerson) *GetCountGender {
	return &GetCountGender{db: db}
}

func (gcg *GetCountGender) Run() ([]int, error) {
	return gcg.db.CountGender()
}