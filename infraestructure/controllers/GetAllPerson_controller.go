package controllers

import (
	"examc1/application"
	"examc1/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllPersonsController struct {
	useCase *application.GetAllPerson
}

func NewGetAllPersonsController() *GetAllPersonsController {

	mysql := infraestructure.GetMySQL()
	useCase := application.NewGetAllPersons(mysql)

	return &GetAllPersonsController{useCase: useCase}

}


func (gap_c *GetAllPersonsController) Run(ctx *gin.Context) {

	persons, err := gap_c.useCase.Run()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Persons": persons})

}