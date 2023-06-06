package account

import (
	"BackendCRM/entities"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r Repository) Save(user *entities.Actors) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	registerData := entities.Register{
		AdminId:      user.ID,
		SuperAdminId: 1,
		Status:       "Inactive",
	}
	return r.db.Create(registerData).Error
}

func (r Repository) FindAll() ([]entities.Actors, error) {
	var actors []entities.Actors
	//err := r.db.Preload("Collections").Order("id").Find(&actors).Error
	err := r.db.Find(&actors).Error
	return actors, err
}

func (r Repository) FindBy(column, value string) (*entities.Actors, error) {
	var actors entities.Actors
	condition := fmt.Sprintf("%s = ?", column)
	// Dapatkan data user dari database berdasarkan ID
	if err := r.db.First(&actors, condition, value).Error; err != nil {
		return nil, err
	}
	return &actors, nil
}

func (r Repository) SoftDel(id string) (*entities.Actors, error) {
	var actors entities.Actors
	// Dapatkan data user dari database berdasarkan ID
	if err := r.db.First(&actors, id).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, err
		}
		return nil, err
	}
	// Hapus data user dari database
	err := r.db.Delete(&actors).Error
	return &actors, err
}

func (r Repository) ChangeById(cst *entities.Actors, id string) (*entities.Actors, error) {
	var existingActor entities.Actors
	// Dapatkan data existingActor dari database berdasarkan ID
	if err := r.db.First(&existingActor, id).Error; err != nil {
		return nil, err
	}
	if cst.Username != "" {
		existingActor.Username = cst.Username
	}
	if cst.Password != "" {
		existingActor.Password = cst.Password
	}
	if cst.RoleId != "" {
		existingActor.RoleId = cst.RoleId
	}
	if cst.Verified != "" {
		existingActor.Verified = cst.Verified
	}
	if cst.Active != "" {
		existingActor.Active = cst.Active
	}
	// Simpan perubahan ke database
	if err := r.db.Save(&existingActor).Error; err != nil {
		return nil, err
	}
	return &existingActor, nil
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
