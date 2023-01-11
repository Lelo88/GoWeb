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
	jsonData := `{"Name":"MacBook Pro","Price":1500,"Published":true}` //como estan puestos las mismas variables no va a tirar error, ya que es estricto

	var p Product

	if err := json.Unmarshal([]byte(jsonData), &p); err!=nil {
		log.Fatal(err)
	}

	fmt.Println(p)
	fmt.Printf("%+v/n",p)
}