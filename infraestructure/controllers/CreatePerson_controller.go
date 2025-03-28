package controllers

import (
	"examc1/application"
	"examc1/domain/entities"
	"examc1/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePersonController struct {
	useCase *application.CreatePerson
}

func NewCreatePersonController() *CreatePersonController {
	
	mysql := infraestructure.GetMySQL()
	useCase := application.NewCreatePerson(mysql)

	return &CreatePersonController{useCase: useCase}

}

func (cp_c *CreatePersonController) Run(ctx *gin.Context) {

	var person entities.Person

	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	if person.Name == "" && person.Sex == "" && person.Gender == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Campos vacíos"})
		return
	}

	err := cp_c.useCase.Run(person)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"Message": "Estudiante creado"})

}

