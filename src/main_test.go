package main

import (
  "testing"
  "io/ioutil"
  // "main"
)

func TestLoadUsersFile(t *testing.T) {
  _, error := ioutil.ReadFile("../test/users.json")
  if error != nil {
    t.Fatalf("failed to load users.json")
  }
}

func TestLoadOrganizationsFile(t *testing.T) {
  _, error := ioutil.ReadFile("../test/organizations.json")
  if error != nil {
    t.Fatalf("failed to load users.json")
  }
}
