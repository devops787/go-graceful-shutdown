package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	fmt.Println("graceful shutdown")

	waitGroup := &sync.WaitGroup{}
	for z := 1; z <= 5; z++ {
		fmt.Printf("starting rountine #%d\n", z)
		waitGroup.Add(1)

		go func(id int, waitGroup *sync.WaitGroup) {
			fmt.Printf("hello from rountine #%d\n", id)

			shutdown := make(chan os.Signal, 1)
			signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
			<-shutdown

			fmt.Printf("exit from rountine #%d\n", id)
			waitGroup.Done()
		}(z, waitGroup)
	}
	waitGroup.Wait()

	fmt.Println("exit from main")
}
