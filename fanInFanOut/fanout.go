package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep() {
	time.Sleep(
		time.Duration(rand.Intn(3000)) * time.Millisecond,
	)
}

func consumer(ch <-chan int, name string) {
	for n := range ch {
		fmt.Printf("Consumer %s <- %d\n", name, n)
	}
}

func producer(ch chan<- int) {
	for {
		sleep()
		n := rand.Intn(100)
		fmt.Printf("-> %d\n", n)
		ch <- n
	}
}

func fanOut(chA <-chan int, chB, chC chan<- int) {

}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA)

	go consumer(chB, "B")
	go consumer(chC, "C")

	fanOut(chA, chB, chC)

}
