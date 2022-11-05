package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fname    string 
	Lname    string 
	Username string 
	Avatar   string 
}
