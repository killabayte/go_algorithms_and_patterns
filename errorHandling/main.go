package main

import "fmt"

type SuperHero struct {
	Name, Role, Email string
	Age               int
}

func (sh SuperHero) Salary() (int, error) {
	switch sh.Role {
	case "Junior problem solver":
		return 500, nil
	case "Middle problem solver":
		return 1000, nil
	case "problem solver":
		return 50000
	case "problem dealer":
		return 500000
	default:
		return 0
	}
}

func (sh *SuperHero) UpdateRole(role string) {
	sh.Role = role
}

func main() {
	sh := SuperHero{
		Name: "Invincible",
		Role: "problem solver",
		Age:  18,
	}

	om := SuperHero{
		Name: "Omni Man",
		Role: "problem dealer",
		Age:  3000,
	}

	fmt.Println(sh)
	fmt.Printf("Super hero says: I'm %s\n", sh.Name)

	fmt.Println(sh.Name, sh.Salary())
	fmt.Println(om.Name, om.Salary())

	sh.UpdateRole("Middle problem solver")
	fmt.Println(sh.Name, sh.Salary())
}
