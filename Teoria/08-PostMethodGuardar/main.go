//vamos a guardar grupos determinados de productos

package main

type request struct { 
	ID 			int `json:"id"`
	Nombre 		string `json:"nombre"`
	Tipo 		string `json:"tipo"`
	Cantidad 	int `json:"cantidad"`
	Precio 		float64 `json:"precio"`
}


func main() {
	
}