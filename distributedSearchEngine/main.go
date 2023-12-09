package main

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
	{Email: "Robbie@example.com", Name: "Robbie"}
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

func main(){
	os.Args[1]
	w := NewWorker(DataBase)
	w.Find(email)
}