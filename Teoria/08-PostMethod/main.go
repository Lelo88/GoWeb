//con este programa realizaremos una peticion post

package main

import "github.com/gin-gonic/gin"

type request struct {
	ID 			int `json:"id"`
	Nombre 		string `json:"nombre"`
	Tipo 		string `json:"tipo"`
	Cantidad 	int `json:"cantidad"`
	Precio 		float64 `json:"precio"`
}


func main() {
	r:= gin.Default()
	r.POST("/products", func(c *gin.Context) { //paso un articulo json sin id, para devolverlo con un id
		var req request
		if err:=c.ShouldBindJSON(&req); err!=nil{ //toma un objeto json y si hay un error, simplemente no hará nada, para que podamos manipular el error
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return 
		}
		req.ID = 4
		c.JSON(200, req) //aca envio el codigo de respuesta en entero acompañado de un objeto
						  //se retorna un json con el id asignado en el caso de que sea correcta la peticion

	})
	r.Run()
}

