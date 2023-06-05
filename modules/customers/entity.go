package customers

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Avatar    string
}

type Collection struct {
	gorm.Model
	FirstName string
	UserID    uint
	Customers Customers
}
