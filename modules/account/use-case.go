package account

import (
	"BackendCRM/entities"
	"BackendCRM/function/hashing"
)

type UseCase struct {
	repo Repository
}

func NewUseCase(repo *repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) Create(user *entities.Actors) error {
	user.Password = hashing.GenerateHash(user.Password)
	return u.repo.Save(user)
}
func (u UseCase) CreateReg(user *entities.Registers) error {
	return u.repo.SaveReg(user)
}
func (u UseCase) Read() ([]entities.Actors, error) {
	return u.repo.FindAll()
}
func (u UseCase) ReadRegis() ([]entities.Registers, error) {
	return u.repo.FindAllRegis()
}

func (u UseCase) ActorReadBy(col, val string) (*entities.Actors, error) {
	return u.repo.ActorFindBy(col, val)
}
func (u UseCase) Delete(id string) (*entities.Actors, error) {
	return u.repo.SoftDel(id)
}

func (u UseCase) Update(cst *entities.Actors, id string) (*entities.Actors, error) {
	return u.repo.ChangeActorById(cst, id)
}
func (u UseCase) UpdateReg(cst *entities.Registers, id string) (*entities.Registers, error) {
	return u.repo.ChangeRegisById(cst, id)
}
