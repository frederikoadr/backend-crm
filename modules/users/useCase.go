package users

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) Create(user *Customers) error {
	return u.repo.Save(user)
}

func (u UseCase) Read() ([]Customers, error) {
	return u.repo.FindAll()
}

func (u UseCase) ReadBy(col, val string) (*Customers, error) {
	return u.repo.FindBy(col, val)
}

func (u UseCase) Delete(id string) (*Customers, error) {
	return u.repo.SoftDel(id)
}

func (u UseCase) Update(cst *Customers, id string) (*Customers, error) {
	return u.repo.ChangeById(cst, id)
}
