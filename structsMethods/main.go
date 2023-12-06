package main

import "fmt"

type SuperHero struct {
	Name, Role, Email string
	Age               int
}

func (sh SuperHero) Salary() int {
	switch sh.Role {
	case "Junior problem solver":
		return 500
	case "Middle problem solver":
		return 1000
	case "problem solver":
		return 50000
	default:
		return 0
	}
}

func main() {
	sh := SuperHero{
		Name: "Invincible",
		Role: "problem solver",
		Age:  18,
	}

	fmt.Println(sh)
	fmt.Printf("Super hero says: I'm %s\n", sh.Name)

	fmt.Println(sh.Salary())
}
