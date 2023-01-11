//ejemplo dado por lucas el 4-1-23 a la tarde, trabajando sobre contextos

package main

import (
	"context"
	"fmt"
	"time"
)

func service(url string, ch chan string) string{
	time.Sleep(time.Second * 3)
	return "data"
}

func fetchData(ctx context.Context, url string) (string, error) {
	ch := make(chan string)
	go service(url, ch)
	for{
		
		select {
			case <-ctx.Done():
				return "", nil
			case data := <-ch:
                return data, nil
    	}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //pasamos un contexto con cancelacion
	
	defer cancel()

	go func() {
		time.Sleep(time.Second)
		cancel()
	}()

	data, err:= fetchData(ctx, "www.google.com")
	if err!= nil {
		fmt.Println("request error", err)
		return
	}

	fmt.Println("request success", data)
}