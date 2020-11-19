package main

import (
	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
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


type Users struct {
	Users []User `json:users`
}

type Organizations struct {
	Organizations []Organization `json:organizations`
}

// NewUsers returns new users
func NewUsers(bytes []byte) (Users, error) {
	var users Users
	json.Unmarshal(bytes, &users)
	return users, nil
}

func isOurJson(t interface{}) bool {
	switch t.(type) {
	case Users:
		return true
	case Organizations:
		return true
	default:
		return false
	}
}

func main() {}
