package main

import (
  "fmt"

  "github.com/pborman/uuid"
)

var currentId int
var users Users

// seed data
func init() {
  RepoCreateUser(User{UUID: uuid.NewUUID(), Name: "John", Email: "john@r.co"})
  RepoCreateUser(User{UUID: uuid.NewUUID(), Name: "Michael", Email: "michael@r.co"})
}

func RepoFindUser(id int) User {
  for _, u := range users {
    if u.Id == id {
      return u
    }
  }

  // return empty user if not found
  return User{}
}

// Ideally, uuid would only get set on user creation & would not be duplicable
func RepoCreateUser(u User) User {
  currentId += 1
  u.Id = currentId
  users = append(users, u)
  return u
}

func DestroyUser(id int) error {
  for i, u := range users {
    if u.Id == id {
      users = append(users[:i], users[i+1:]...)
      return nil
    }
  }

  return fmt.Errorf("Could not find User with id of %d to delete", id)
}
