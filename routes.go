package main

import "net/http"

type Route struct {
  Name          string
  Method        string
  Pattern       string
  HandlerFunc   http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "Users",
    "GET",
    "/",
    UserShow,
  },
  Route{
    "User",
    "GET",
    "/{userId}",
    UserShow,
  },
  Route{
    "UserCreate",
    "POST",
    "/",
    UserCreate,
  },
  Route{
    "User",
    "PUT",
    "/{userId}",
    UserUpdate,
  },
  Route{
    "User",
    "DELETE",
    "/{userId}",
    UserDelete,
  },
}
