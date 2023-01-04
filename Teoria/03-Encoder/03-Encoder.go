package main

import (
	"encoding/json"
	"os"
)

type MyData struct {
	ProductoID  string
	Precio		float64	
}

func main(){
	myEncoder := json.NewEncoder(os.Stdout)
	data := MyData {
		ProductoID: "XSAD",
		Precio: 25.5,
	}

	myEncoder.Encode(data)
}