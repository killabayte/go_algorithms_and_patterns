package main

import "fmt"

func main() {
	//Slices
	n := []int{1, 3, 5, 7, 9}
	fmt.Println(n[len(n)-1])
	fmt.Println(append(n, 11, 13, 15, 17, 19))
}
