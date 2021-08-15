package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		c <- "foo"
	}()
	/*
		The first goroutine is blocked after sending the message `foo` since no
		receivers are yet ready.

		" If the capacity is zero or absent, the channel is unbuffered and
		communication succeeds only when both a sender and receiver are ready. "
	*/

	go func() {
		defer wg.Done()

		time.Sleep(time.Second * 1)
		fmt.Println(`Message: ` + <-c)
	}()

	wg.Wait()
}
