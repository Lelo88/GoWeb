package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Name string 
	Price int 
	Published bool
}

func main(){
	p:= Product {
		Name: "MacBook Pro",
		Price: 1500,
		Published: true,
	}

	jsonData, err := json.Marshal(p) //el objeto p se transforma en un objeto json
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}