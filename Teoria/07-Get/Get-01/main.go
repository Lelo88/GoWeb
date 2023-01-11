package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string `json:"message"`
	Data any	
}

func main() {
	sv := gin.Default()
	sv.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	}) //al escribit en la url me devuelve pong

	sv.GET("/next/:number", func(c *gin.Context) {
		n, err := strconv.Atoi(c.Param("number"))
		if err!= nil { 
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "failed to parse number", 
				"data": nil,
			})
			return
        }
        n++

		c.JSON(http.StatusOK, gin.H{
			"message": "status OK",
			"data": n,
		}) 
	}) //al escribit en la url un numero, me devuelve el siguiente, sino me devuelve un error

    sv.Run()
}

//todo hay que probarlo en el postman