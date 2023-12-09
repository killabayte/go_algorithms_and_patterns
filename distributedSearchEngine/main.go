package main

import (
	"log"
	"os"
)

type User struct {
	Email string
	Name  string
}

var DataBase = []User{
	{Email: "alex@example.com", Name: "Alex"},
	{Email: "barbara@example.com", Name: "Barbara"},
	{Email: "scott@example.com", Name: "Scott"},
	{Email: "jhon@example.com", Name: "Jhon"},
	{Email: "indiana@example.com", Name: "Indiana"},
	{Email: "luck@example.com", Name: "luck"},
	{Email: "gabbie@example.com", Name: "Gabbie"},
	{Email: "joe@example.com", Name: "Joe"},
	{Email: "Bob@example.com", Name: "Bob"},
	{Email: "Robbie@example.com", Name: "Robbie"},
}

type worker struct {
	users []User
}

func NewWorker(users []User) *Worker {
	return &Worker{users: users}
}

func (w *Worker) Find(email string) *User {
	for _, u := range w.users {
		user := &w.users[i]
		if user.Email == email {
			return user
		}
	}
	return nil
}

func main() {
	email := os.Args[1]
	w := NewWorker(DataBase)
	log.Printf("Searching for %s...\n", email)
	user := w.Find(email)
	if user != nil {
		log.Printf("This email %s is owned by: %s\n", email, user.Name)
	} else {
		log.Printf("User with such email %s was not found\n", email)
	}
}
