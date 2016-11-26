package main

import (
  "time"

  "github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  Name              string          `json:"name"`
  Email             string          `json:"email"`
  HashedPassword    string          `json:"hashed_password"`
  Salt              string          `json:"salt"`
  AppIds            []AppId         `json:"app_ids"`            `gorm:"many2many:user_apps;"`
  Permissions       []Permission    `json:"permissions"`        `gorm:"many2many:user_permissions"`
  JobTitle          string          `json:"job_title"`
  Photo             string          `json:"photo_file_name"`
  Bio               string          `json:"bio"`
}

type Users []User

type AppdId struct {
  gorm.Model
  Name        string
}

type Permission struct {
  gorm.Model
  Level       string
}

// Make this happen somehow
db.Model(&user).Related(&app_ids, "AppIds")
db.Model(&user).Related(&permissions, "Permissions")
