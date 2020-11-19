package main

import (
  "encoding/json"
  "fmt"
  "os"
  "io/ioutil"
)

type User struct {
  Id int
  Name string
}

type Organization struct {
  Uid int
  Uname string
  User []User
}
