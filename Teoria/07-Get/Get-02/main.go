//vamos a simular traer y llevar una base de datos para realizar las proximas ejercitaciones

package main

import "github.com/gin-gonic/gin"

type Websites struct {
	URL 		string `json:"url"`
	Host 		string `json:"host"`
	Category 	string `json:"category"`
	Protected 	bool `json:"protected"`
}

var websites = []Websites{ //pseudo base de datos creada a traves de un slice con la estructura Websites de arriba
	{URL: "www.google.com", Host: "google", Category: "search", Protected: false},
	{URL: "www.yahoo.com", Host: "yahoo", Category: "search", Protected: true},
	{URL: "www.mercadolibre.com", Host: "meli", Category: "e-commerce", Protected: false}, 
}

//vamos a obtener toda la info de la base de datos para pasarla como json a traves de postman
func Get() []Websites{
	return websites
}

func GetQuery(category string) []Websites {
	websites := Get() //pasamos todos los datos de nuestra base de datos

	var filtered []Websites
	for _, website := range websites {
        if category != "" && (website.Category != category) {
			continue
		}
		filtered = append(filtered, website)
	}
	return filtered
}

func main() {
	sv:= gin.Default()	
	sv.GET("/websites", func(c *gin.Context) {
		//request ... no tenemos nada porque no pasamos nada por el body
			

		//buscamos por parametro un dato de nuestra base de datos. lo hacemos con querys
		queryCategory := c.Query("category") //en la web anotamos websites?category= tipo de categoria
		websites2 := GetQuery(queryCategory)
		c.JSON(200, gin.H{
			"message": "success",
			"data": websites2,
		})	
	})

	//mostramos toda la base de datos
	sv.GET("/webs", func(c *gin.Context) {
		websites = Get()
		c.JSON(200, gin.H{
			"message": "success",
			"data": websites,
		})
	})

	sv.Run()
}