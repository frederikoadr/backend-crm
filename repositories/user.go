package repositories

import (
	"BackendCRM/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	GetByID(id int) []entities.User
}

func (repo UserRepository) GetByID(id int) []entities.User {
	// implementasi query get user by id
	return []entities.User{}
}
