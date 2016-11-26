package main

import (
  "encoding/json"
  "net/http"
  "strconv"
  "io"
  "io/ioutil"

  "github.com/gorilla/mux"
)

func UserShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  if uid, ok := vars["userId"]; ok {
    // single user view
    userId, _ := strconv.Atoi(uid)

    user := RepoFindUser(userId)

    if user.Id > 0 {
      w.Header().Set("Content-Type", "application/json; charset=UTF-8")
      w.WriteHeader(http.StatusOK)

      if err := json.NewEncoder(w).Encode(user); err != nil {
          panic(err)
      }

      return
    }

    // If we didn't find it, 404
  	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  	w.WriteHeader(http.StatusNotFound)

  	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
  		panic(err)
  	}
  } else {
    // index view
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(users); err != nil {
        panic(err)
    }
  }
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
  var user User
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

  // Error handling
  if err != nil {
    panic(err)
  }
  if err := r.Body.Close(); err != nil {
    panic(err)
  }
  if err := json.Unmarshal(body, &user); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422) // unprocessabile entity

    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  u := RepoCreateUser(user)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)

  if err := json.NewEncoder(w).Encode(u); err != nil {
    panic(err)
  }
}
