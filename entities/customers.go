package entities

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Avatar    string
}
