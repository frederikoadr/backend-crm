package account

import (
	"BackendCRM/entities"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r repository) Save(user *entities.Actors) error {
	return r.db.Create(user).Error
}
func (r repository) SaveReg(user *entities.Registers) error {
	return r.db.Create(user).Error
}
func (r repository) FindAll() ([]entities.Actors, error) {
	var actors []entities.Actors
	//err := r.db.Preload("Collections").Order("id").Find(&actors).Error
	err := r.db.Find(&actors).Error
	return actors, err
}
func (r repository) FindAllRegis() ([]entities.Registers, error) {
	var registers []entities.Registers
	//err := r.db.Preload("Collections").Order("id").Find(&registers).Error
	err := r.db.Find(&registers).Error
	return registers, err
}

func (r repository) ActorFindBy(column, value string) (*entities.Actors, error) {
	var actors entities.Actors
	condition := fmt.Sprintf("%s = ?", column)
	// Dapatkan data user dari database berdasarkan ID
	if err := r.db.First(&actors, condition, value).Error; err != nil {
		return nil, err
	}
	return &actors, nil
}

func (r repository) SoftDel(id string) (*entities.Actors, error) {
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

func (r repository) ChangeActorById(cst *entities.Actors, id string) (*entities.Actors, error) {
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
func (r repository) ChangeRegisById(cst *entities.Registers, id string) (*entities.Registers, error) {
	var existingReg entities.Registers
	// Dapatkan data existingReg dari database berdasarkan ID
	if err := r.db.First(&existingReg, id).Error; err != nil {
		return nil, err
	}
	fmt.Println(existingReg.AdminId)
	if cst.Status != "" {
		existingReg.Status = cst.Status
	}
	// Simpan perubahan ke database
	if err := r.db.Save(&existingReg).Error; err != nil {
		return nil, err
	}
	return &existingReg, nil
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

type Repository interface {
	Save(user *entities.Actors) error
	SaveReg(user *entities.Registers) error
	FindAll() ([]entities.Actors, error)
	FindAllRegis() ([]entities.Registers, error)
	ActorFindBy(column, value string) (*entities.Actors, error)
	SoftDel(id string) (*entities.Actors, error)
	ChangeActorById(cst *entities.Actors, id string) (*entities.Actors, error)
	ChangeRegisById(cst *entities.Registers, id string) (*entities.Registers, error)
}
