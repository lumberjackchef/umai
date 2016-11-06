package main

import (
  "encoding/json"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
  "github.com/pborman/uuid"
)

var users Users = Users{
  User{UUID: uuid.NewUUID(), Id: 1, Name: "John", Email: "john@r.co"},
  User{UUID: uuid.NewUUID(), Id: 2, Name: "Michael", Email: "michael@r.co"},
}

// Ideally, uuid would only get set on user creation & would not be duplicable

func UserShow(w http.ResponseWriter, r *http.Request) {
  var user User
  vars := mux.Vars(r)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if uid, ok := vars["userId"]; ok {
    // single user view
    userId, _ := strconv.Atoi(uid)

    // user the search algorithm here in the future, it's more efficient
    for _, u := range users {
      if u.Id == userId {
        user = u
      }
    }

    if err := json.NewEncoder(w).Encode(user); err != nil {
        panic(err)
    }
  } else {
    // index view
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }
  }
}
