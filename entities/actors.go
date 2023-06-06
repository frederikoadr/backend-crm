package entities

import "gorm.io/gorm"

type Actors struct {
	Username string
	Password string
	RoleId   string
	Verified string
	Active   string
	gorm.Model
}
