package main

func main() {
	in := make(chan int)
	out := make(chan int)

	go echoWorker(in, out)
}
