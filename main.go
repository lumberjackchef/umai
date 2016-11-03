package main

import (
    "log"
    "net/http"
    "strconv"

    // "time"
    "encoding/json"

    "github.com/gorilla/mux"
    "github.com/pborman/uuid"
)


//----------//
// need to abstract out all this users crap to a db

type User struct {
  UUID              uuid.UUID   `json:"uuid"`
  Id                int         `json:"id"`
  Name              string      `json:"name"`
  Email             string      `json:"email"`
  // Hashed_password   string      ``                        // omit these somehow
  // Salt              string      ``                        // omit these somehow
  // App_ids           []int       `json:"app_ids"`
  // Created_at        time.Time   `json:"created"`
  // Updated_at        time.Time   `json:"updated"`
  // permissions       []string    `json:"permissions"`
  // Job_title         string      `json:"job_title"`
  // Photo             string      `json:"photo_file_name"`
  // Bio               string      `json:"bio"`
}

type Users []User

var users Users = Users{
  User{UUID: uuid.NewUUID(), Id: 1, Name: "John", Email: "john@r.co"},
  User{UUID: uuid.NewUUID(), Id: 2, Name: "Michael", Email: "michael@r.co"},
}

//----------//



func main() {

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", UserShow)
  router.HandleFunc("/{userId}", UserShow) // need a way to force correct pattern here

  log.Fatal(http.ListenAndServe(":8000", router))
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

    // user the search algorythm here in the future, it's more efficient
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
