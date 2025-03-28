package main

import (
	"examc1/infraestructure"
	"examc1/infraestructure/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Iniciamos las dependencias
	infraestructure.InitDependencies()

	cors.Default()

	routes.RegisterRoutes(r)

	// Levantamos el servidor
	r.Run()

}