package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)

	ch <- 10
	ch <- 20
	ch <- 30
	ch <- 40

	go func() {
		for {
			msg := <-ch
			fmt.Println(msg)
		}
	}()

	time.Sleep(1000 * time.Microsecond)
}
