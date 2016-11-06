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

// Ideally, uuid would only get set on user creation & would not be duplicable

//----------//



func main() {

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", UserShow)
  router.HandleFunc("/{userId}", UserShow)

  log.Fatal(http.ListenAndServe(":8000", router))
}

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
