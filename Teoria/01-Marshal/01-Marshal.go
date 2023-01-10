package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Name string `json:"name"`	//notacion snake para formato json
	Price int 	`json:"price"`	
	Published bool `json:"is_published"`	
}

func main(){
	p:= Product {
		Name: "MacBook Pro",
		Price: 1500,
		Published: true,
	}

	jsonData, err := json.Marshal(p) //el objeto p se transforma en un objeto json y lo imprime por consola
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}