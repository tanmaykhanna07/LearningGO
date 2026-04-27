package main

import (
	"fmt"
	"time"
)

func FetchFast(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Fast API"
}
func FetchSlow(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Slow API"
}
func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)
	go FetchFast(channel1)
	go FetchSlow(channel2)
	select {
	case result := <-channel1:
		fmt.Println("Fast finished first: ", result)

	case result := <-channel2:
		fmt.Println("Fast finished first: ", result)
	}
}
