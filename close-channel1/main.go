package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(ch)			// el emisor cierra (señal para decir que we are done, homs)
	}()

	// for valor := range ch {
	// 	fmt.Println("Recibi: ", valor)
	// }

	for {
		valor, isOpen := <- ch 
		if !isOpen {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Value:", valor)
	}

	fmt.Println("Canal cerrado. Fin del programa.")
}
