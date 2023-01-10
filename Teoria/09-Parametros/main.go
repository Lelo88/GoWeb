//los parametros son los datos que pasamos por enlace para que el servidor los lea por URL
//Los Query params son los caracteres que se utilizan en el uso de parametros en la URL

package main

import "github.com/gin-gonic/gin"

//definimos una pseudobase de datos
var empleados = map[string]string {
	"644" : "Empleado A",
	"755" : "Empleado B",
	"777" : "Empleado C",
}

func PaginaPrincipal(c *gin.Context) {
	c.String(200,"¡Bienvenido a la empresa Gophers!")
}

func BuscarEmpleado(c *gin.Context) {
	empleado, ok := empleados[c.Param("id")]

	if ok {
        c.String(200,"Informacion del empleado %s, nombre, %s", c.Param("id"), empleado)
	} else {
		c.String(404, "Empleado %s no encontrado", c.Param("id"))
    }	
}

func main() {
	server := gin.Default()	
	server.GET("/", PaginaPrincipal)
	server.GET("/empleado/:id", BuscarEmpleado)
	server.Run(":8080")
}

/*cuando empiezo a correr el programa, me dirijo a postman y coloco el metodo get con el enlace
luego de eso, pruebo pasando los parametros para ver si existe en mi base de datos ficticia*/