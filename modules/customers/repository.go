package customers

import (
	"BackendCRM/entities"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r repository) Save(user *entities.Customers) error {
	return r.db.Create(user).Error
}

func (r repository) FindAll() ([]entities.Customers, error) {
	var customers []entities.Customers
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r repository) FindBy(column, value string) (*entities.Customers, error) {
	var customers entities.Customers
	condition := fmt.Sprintf("%s = ?", column)
	// Dapatkan data user dari database berdasarkan ID
	if err := r.db.First(&customers, condition, value).Error; err != nil {
		return nil, err
	}
	return &customers, nil
}

func (r repository) SoftDel(id string) (*entities.Customers, error) {
	var customers entities.Customers
	// Dapatkan data user dari database berdasarkan ID
	if err := r.db.First(&customers, id).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, err
		}
		return nil, err
	}
	// Hapus data user dari database
	err := r.db.Delete(&customers).Error
	return &customers, err
}

func (r repository) ChangeById(cst *entities.Customers, id string) (*entities.Customers, error) {
	var existingCustomer entities.Customers
	// Dapatkan data existingCustomer dari database berdasarkan ID
	if err := r.db.First(&existingCustomer, id).Error; err != nil {
		return nil, err
	}
	if cst.FirstName != "" {
		existingCustomer.FirstName = cst.FirstName
	}
	if cst.LastName != "" {
		existingCustomer.LastName = cst.LastName
	}
	if cst.Email != "" {
		existingCustomer.Email = cst.Email
	}
	if cst.Avatar != "" {
		existingCustomer.Avatar = cst.Avatar
	}

	// Simpan perubahan ke database
	if err := r.db.Save(&existingCustomer).Error; err != nil {
		return nil, err
	}
	return &existingCustomer, nil
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

type Repository interface {
	Save(user *entities.Customers) error
	FindAll() ([]entities.Customers, error)
	FindBy(column, value string) (*entities.Customers, error)
	SoftDel(id string) (*entities.Customers, error)
	ChangeById(cst *entities.Customers, id string) (*entities.Customers, error)
}
