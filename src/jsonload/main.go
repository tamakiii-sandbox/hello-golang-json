package jsonload

import (
	"encoding/json"
	// "fmt"
	// "os"
	// "io/ioutil"
)

// User is user
type User struct {
	Id   int
	Name string
}

// Organization is organization
type Organization struct {
	Uid   int
	Uname string
	User  []User
}

// NewUsers returns new users
func NewUsers(bytes []byte) ([]User, error) {
	var users []User
	json.Unmarshal(bytes, &users)
	return users, nil
}

func main() {}
