package main

import (
	"context"
	"fmt"
)

func addValue (ctx context.Context) context.Context {
	return context.WithValue(ctx, "bootcamp", "Go")
}

func main(){
	ctx := context.Background() //funcion para inicializar un contexto vacio
	newCtx := addValue(ctx)

	variable := newCtx.Value("bootcamp")

	fmt.Println(variable)

}

