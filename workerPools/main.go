package main

import (
	"math/rand"
	"time"
)

func echoWorker(in, out chan int) {
	for {
		n := in
		time.Sleep(
			time.Duration(rand.Intn(3000)) * time.Millisecond,
		)
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go echoWorker(in, out)
}
