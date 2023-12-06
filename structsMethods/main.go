package main

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

}
