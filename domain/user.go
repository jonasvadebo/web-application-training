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

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	//This could be done above
	nextID++
	users = append(users, &u)
	return u, nil
}
