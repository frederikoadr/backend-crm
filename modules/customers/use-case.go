package customers

import "BackendCRM/entities"

type UseCase struct {
	repo Repository
}

func NewUseCase(repo *repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) Create(user *entities.Customers) error {
	return u.repo.Save(user)
}

func (u UseCase) Read() ([]entities.Customers, error) {
	return u.repo.FindAll()
}

func (u UseCase) ReadBy(col, val string) (*entities.Customers, error) {
	return u.repo.FindBy(col, val)
}

func (u UseCase) Delete(id string) (*entities.Customers, error) {
	return u.repo.SoftDel(id)
}

func (u UseCase) Update(cst *entities.Customers, id string) (*entities.Customers, error) {
	return u.repo.ChangeById(cst, id)
}
