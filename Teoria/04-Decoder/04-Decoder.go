package main

import (
	"encoding/json"
	"io"
	"log"
	"strings"
)

const jsonStream = `
{"ProductoID":"REWR4","Precio":25.5}
{"ProductoID":"4FSDF","Precio":28.4}
{"ProductoID":"7HGHG","Precio":22.3}
` //creamos un dato string similar a un json

type MyData struct { //creamos un struct para pasar los datos json a este
	ProductoID  string
	Precio		float64	
}

func main(){
	myStreaming := strings.NewReader(jsonStream) //variable que leera la constante
	myDecoder := json.NewDecoder(myStreaming) //variable que decodificara la constante en un archivo json
	
	for {
		var data MyData
		if err := myDecoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} 
		log.Println("Datos recibidos: ", data.ProductoID, data.Precio)
	}
}