package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"

    "time"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/pborman/uuid"
)


// need to abstract out all this users crap to a db
// << EOF

type User struct {
  UUID      uuid.UUID   `json:"uuid"`
  Id        int         `json:"id"`
  Name      string      `json:"name"`
  Email     string      `json:"email"`
  Created   time.Time   `json:"created"`
}

type Users []User

var users Users = Users{
  User{UUID: uuid.NewUUID(), Id: 1, Name: "John", Email: "john@r.co"},
  User{UUID: uuid.NewUUID(), Id: 2, Name: "Michael", Email: "michael@r.co"},
}

// EOF



func main() {

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)
  router.HandleFunc("/users", UserShow)
  router.HandleFunc("/users/{userId}", UserShow) // need a way to force correct pattern here

  log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "KHAAAAAAAAAAAAAN!")
}

func UserShow(w http.ResponseWriter, r *http.Request) {
  var response User
  vars := mux.Vars(r)
  _, ok := vars["userId"]

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if ok {
    // single user view
    userId, _ := strconv.Atoi(vars["userId"])

    for _, p := range users {
      if p.Id == userId {
        response = p
      }
    }

    if err := json.NewEncoder(w).Encode(response); err != nil {
        panic(err)
    }
  } else {
    // index view
    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }
  }
}
