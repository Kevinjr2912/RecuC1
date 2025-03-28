package controllers

import (
	"examc1/application"
	"examc1/infraestructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetCountGendersController struct {
	useCase *application.GetCountGender
}

func NewGetCountGendersController() *GetAllPersonsController {
	
	mysql := infraestructure.GetMySQL()
	useCase := application.NewGetCountGender(mysql)

	return &GetAllPersonsController{useCase: (*application.GetAllPerson)(useCase)}

}

func (gcg_c *GetCountGendersController) Run(ctx *gin.Context) {

	values, err := gcg_c.useCase.Run()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	countH := values[0]
	countF := values[1]


	ctx.JSON(http.StatusOK, gin.H{
		"Hombres": countH,
		"Mujeres": countF,
	})


}	