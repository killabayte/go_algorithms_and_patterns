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
	{Email: "bob@example.com", Name: "Bob"},
	{Email: "bobbie@example.com", Name: "Robbie"},
}

type Worker struct {
	users []User
	ch    chan *User
}

func NewWorker(users []User, ch chan *User) *Worker {
	return &Worker{users: users, ch: ch}
}

func (w *Worker) Find(email string) {
	for i := range w.users {
		user := &w.users[i]
		if user.Email == email {
			w.ch <- user
		}
	}
}

func main() {
	email := os.Args[1]
	ch := make(chan *User)
	w := NewWorker(DataBase, ch)
	log.Printf("Searching for %s...\n", email)
	go w.Find(email)
	if user != nil {
		log.Printf("This email %s is owned by: %s\n", email, user.Name)
	} else {
		log.Printf("User with: %s, 'email' was not found\n", email)
	}
}
