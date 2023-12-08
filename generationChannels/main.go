package main

import (
	"math/rand"
	"time"
)

func process2(ch chan int) {
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
}

func main() {
	ch := make(chan int)
	go process2()
}
