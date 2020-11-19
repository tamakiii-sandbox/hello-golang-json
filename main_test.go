package main

import (
	"io/ioutil"
	"testing"
	"fmt"
	// "reflect"
)

func TestLoadUsersFile(t *testing.T) {
	_, error := ioutil.ReadFile("./testdata/users.json")
	if error != nil {
		t.Fatalf("failed to load users.json")
	}
}

func TestLoadOrganizationsFile(t *testing.T) {
	_, error := ioutil.ReadFile("./testdata/organizations.json")
	if error != nil {
		t.Fatalf("failed to load users.json")
	}
}

func TestUnmarshalUsers(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./testdata/users.json")
	users, error := NewUsers(bytes)

	if error != nil {
		t.Fatalf("failed to new users")
	}

	fmt.Printf("%v", users)
	fmt.Printf("%v", isOurJson(users))
}

