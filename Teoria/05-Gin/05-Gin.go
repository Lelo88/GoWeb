package main

import "github.com/gin-gonic/gin"

func main(){
	//se crea un router con gin
	router := gin.Default()

	// Captura la solicitud "Hello World"
	router.GET("/hello-world", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//corremos nuestro servidor sobre el puerto 8080
	router.Run()
}