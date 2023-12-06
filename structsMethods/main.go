package main

import "fmt"

type SuperHero struct {
	Name, Role, Email string
	Age               int
}

func main() {
	sh := SuperHero{
		Name: "Invincible",
		Role: "Problem Solver",
		Age:  18,
	}

	fmt.Println(sh)
	fmt.Printf("Super hero says: I'm %s\n", sh.Name)
}
