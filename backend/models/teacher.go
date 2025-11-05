package models

import "gorm.io/gorm"

type Teacher struct {
    gorm.Model
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}
// /