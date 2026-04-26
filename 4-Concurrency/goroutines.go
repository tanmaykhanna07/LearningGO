package main

import (
	"fmt"
	"sync"
	"time"
)

func SlowTask(wg *sync.WaitGroup) {
	fmt.Println("Working... ")
	time.Sleep(2 * time.Second)
	fmt.Println("Finished slow task")
	wg.Done()
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go SlowTask(&waitGroup)
	waitGroup.Wait()
	fmt.Println("Finished main task")
}
