package jsonload

import (
	"io/ioutil"
	"jsonload"
	"testing"
)

func TestLoadUsersFile(t *testing.T) {
	_, error := ioutil.ReadFile("testdata/users.json")
	if error != nil {
		t.Fatalf("failed to load users.json")
	}
}

func TestLoadOrganizationsFile(t *testing.T) {
	_, error := ioutil.ReadFile("testdata/organizations.json")
	if error != nil {
		t.Fatalf("failed to load users.json")
	}
}

func TestUnmarshalUsers(t *testing.T) {
	bytes, _ := ioutil.ReadFile("testdata/users.json")
	_, error := jsonload.NewUsers(bytes)

	if error != nil {
		t.Fatalf("failed to new users")
	}
}
