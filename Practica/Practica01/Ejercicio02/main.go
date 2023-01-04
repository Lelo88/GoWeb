package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persona struct {
	Nombre 		string `json:"nombre"`
	Apellido 	string `json:"apellido"`
}

func main(){
	
	router := gin.Default()

	router.POST("/saludo", func(c *gin.Context){
		var p Persona

		c.Bind(&p)

		saludo := p.Nombre + " " + p.Apellido
		c.String(http.StatusOK, "Hola " + saludo)
	})
	
	router.Run(":8080")
}