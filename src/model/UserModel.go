package model

type User struct {
	Name     string
	LastName string
	Age      int
}

func CreateNicolas() *User {

	return &User{"Nicolas", "canicatti ", 28}
}

func CreateUser(name string, lastName string, age int) *User {

	return &User{name, lastName, age}
}
