package main

import (
	"log"
	"time"
)

func startConsumers(queue chan int, name string) {
	for {
		time.Sleep(time.Second)
		log.Println(name, <-queue)
	}
}

func main() {
	n := 10
	queue := make(chan int, n)

	for i := 1; i <= n; i++ {
		queue <- i
	}

	go startConsumers(queue, "C1")
	go startConsumers(queue, "C2")

	time.Sleep(time.Second * 10)
}
