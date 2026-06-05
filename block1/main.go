package main

import (
	"fmt"
	//"time"
 

func main() {
	ch := make(chan string)			// without buffer
	go func() {
		fmt.Println("Goroutine: voy a enviar en 1 segundo")
		//time.Sleep(1 * time.Second)
		ch <- "Holaaaa"
		fmt.Println("Goroutine: done sending ...")
	}()

	fmt.Println("Main: esperando recibir ...")
	msg := <- ch 
	fmt.Println(msg)
}
