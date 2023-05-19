package domain

type User struct {
	ID        int
	FirstName string
	LastName  string
}

// Package declared variables.
var (
	//Declare variables
	//Slice of pointers to users
	users  []*User
	nextID = 1
)
