package main

import (
  // "time"

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
