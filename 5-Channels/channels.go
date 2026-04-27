package main

import (
	"fmt"
	"time"
)

func FetchPrice(Ticker string, channel chan float64) {
	time.Sleep(2 * time.Second)
	channel <- 99.99
}

func main() {
	ch := make(chan float64)
	go FetchPrice("AAPL", ch)
	result := <-ch
	fmt.Println(result)
}
