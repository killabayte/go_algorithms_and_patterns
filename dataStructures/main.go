package main

import "fmt"

func main() {
	//Slices - ordered
	n := []int{1, 3, 5, 7, 9}
	fmt.Println(n[len(n)-1])
	fmt.Println(append(n, 11, 13, 15, 17, 19))

	//Maps - non ordered
	sh := map[string]string{
		"name":  "Invinceble",
		"email": "invincible@amazon.com",
		"role":  "problem solver",
	}
	fmt.Println(sh)
	age, ok := sh["age"]
	if !ok {
		fmt.Println("Age not found")
	} else {
		fmt.Println("Age:", age)
	}

}
