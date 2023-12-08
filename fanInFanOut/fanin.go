package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep(){
	time.Sleep(
		time.Duration(rand.Intn(3000)) * time.Millisecond
	)
}

func producer(ch chan<- int, name string) {
	for {
		sleep()
		n := rand.Intn(100)
		fmt.Printf("Channel %s -> %d\n", name, n)
		ch <- n
	}
}

func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)
}
