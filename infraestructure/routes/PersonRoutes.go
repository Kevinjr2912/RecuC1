package routes

import (
	"examc1/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	routes := router.Group("/persons")
	{
		routes.POST("/", controllers.NewCreatePersonController().Run)
		routes.GET("/newPersonIsAdded", controllers.NewGetAllPersonsController().Run)
		routes.GET("/countGender", controllers.NewGetCountGendersController().Run)
	}

}