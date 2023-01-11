package main

import (
	"encoding/json"
	"os"
)

type MyData struct {
	ProductoID  string `json:"producto_id"`
	Precio		float64	`json:"precio"`
}

func main(){
	myEncoder := json.NewEncoder(os.Stdout)
	data := MyData {
		ProductoID: "XSAD",
		Precio: 25.5,
	}

	myEncoder.Encode(data)
}