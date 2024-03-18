package user

import (
	"fmt"
	"log"
)

type User struct {
	Name string
	Age  int
}

func (u User) CreateUser() bool {
	log.Printf("User created")
	return true
}

func (u User) DeleteUser() bool {
	log.Printf("User deleted")
	return true
}

func (u User) UpdateUser(name string, age string) bool {
	log.Printf("User updated")
	return true
}

func (u User) GetUsers() {
	log.Printf("User name", u.Name, " age : ", u.Age)
}

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Greetings(name string) {
	log.Printf("Hello %v, this is the updated version", name)
}
