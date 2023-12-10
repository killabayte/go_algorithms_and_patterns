package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var (
	timeInput = os.Args[1]
)

func main() {
	// Your timestamp in milliseconds
	milliseconds, err := strconv.ParseInt(timeInput, 10, 64)
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	// Convert to seconds by dividing by 1000
	seconds := milliseconds / 1000

	// Convert to time.Time
	t := time.Unix(seconds, 0)

	// Print the formatted time
	fmt.Println("Formatted Time:", t)
}
