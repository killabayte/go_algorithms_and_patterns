package main

// Switch is the second one
func price(item string) int {
	switch item {
	case "apple":
		return 10
	case "orange":
		return 20
	case "carrot":
		return 5
	default:
		return 0
	}

}

func main() {
	//For loop - one of the control structure
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			println(i)
		}
	}

	price("apple")
}
