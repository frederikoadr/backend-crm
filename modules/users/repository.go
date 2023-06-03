package users

import (
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) Save(user *Customers) error {
	return r.db.Create(user).Error
}

func (r Repository) FindAll() ([]Customers, error) {
	var customers []Customers
	//err := r.db.Preload("Collections").Order("id").Find(&customers).Error
	err := r.db.Find(&customers).Error
	return customers, err
}

func (r Repository) SoftDel(id string) (*Customers, error) {
	var customers Customers
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

func (r Repository) ChangeById(cst *Customers, id string) (*Customers, error) {
	var existingCustomer Customers
	// Dapatkan data existingCustomer dari database berdasarkan ID
	if err := r.db.First(&existingCustomer, id).Error; err != nil {
		return nil, err
	}
	existingCustomer.FirstName = cst.FirstName
	existingCustomer.LastName = cst.LastName
	existingCustomer.Email = cst.Email
	existingCustomer.Avatar = cst.Avatar
	// Simpan perubahan ke database
	if err := r.db.Save(&existingCustomer).Error; err != nil {
		return nil, err
	}
	return &existingCustomer, nil
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
