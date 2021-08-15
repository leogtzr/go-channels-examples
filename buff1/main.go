package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// capacity = 2
	c := make(chan string, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()

		c <- `foo`
		c <- `bar`
	}()

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 1)

		fmt.Println(`Msg: ` + <-c)
		fmt.Println(`Msg 2: ` + <-c)
	}()

	wg.Wait()
}
