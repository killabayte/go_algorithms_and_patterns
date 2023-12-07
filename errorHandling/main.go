package main

import (
	"errors"
	"fmt"
)

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
		return 50000, nil
	case "problem dealer":
		return 500000, nil
	default:
		return 0, errors.New("Error to handle this role")
	}
}

func (sh *SuperHero) UpdateRole(role string) {
	sh.Role = role
}

func main() {
	user := SuperHero{Name: "Robot", Role: "Architecture Designer"}
	salary, err := user.Salary()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Salary:", salary)
	}
}
