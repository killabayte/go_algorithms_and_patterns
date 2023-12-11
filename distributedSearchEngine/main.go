package main

import (
	"log"
	"os"
	"time"
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

func NewWorker(users []User, ch chan *User, name string) *Worker {
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

	// Split the database into four parts
	partSize := len(DataBase) / 4
	firstPart := DataBase[:partSize]
	secondPart := DataBase[partSize : 2*partSize]
	thirdPart := DataBase[2*partSize : 3*partSize]
	fourthPart := DataBase[3*partSize:]

	log.Printf("Searching for %s...\n", email)
	go NewWorker(firstPart, ch, "#1").Find(email)
	go NewWorker(secondPart, ch, "#2").Find(email)
	go NewWorker(thirdPart, ch, "#3").Find(email)
	go NewWorker(fourthPart, ch, "#4").Find(email)
	go w.Find(email)

	select {
	case user := <-ch:
		log.Printf("This email %s is owned by: %s\n", email, user.Name)
	case <-time.After(100 * time.Millisecond):
		log.Printf("The email %s was not found\n", email)
	}
}
