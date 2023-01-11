package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	Name string `json:"name"`	//notacion snake para formato json
	Price int 	`json:"price"`	
	Published bool `json:"is_published,omitempty"`	
}

func main(){
	p:= Product {
		Name: "MacBook Pro",
		Price: 1500,
	}

	//jsonData, err := json.MarshalIndent(p,"","  ") //el objeto p se transforma en un objeto json y lo imprime por consola // lo podemos identar
	jsonData, err := json.Marshal(p) // -> este es el que me conviene
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}