package main

import (
	"context"
	"fmt"
)

var programacion = map[string]string {"bootcamp":"go"}

func addValue (ctx context.Context) context.Context {
	return context.WithValue(ctx, programacion, "go")
}

func main(){
	ctx := context.Background() //funcion para inicializar un contexto vacio
	newCtx := addValue(ctx)

	variable := newCtx.Value("bootcamp")

	fmt.Println(variable)


}

