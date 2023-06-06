package account

import (
	"BackendCRM/entities"
	"crypto/sha256"
	"fmt"
)

type UseCase struct {
	repo *Repository
}

func NewUseCase(repo *Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (u UseCase) Create(user *entities.Actors) error {
	user.Password = generateHash(user.Password)
	return u.repo.Save(user)
}

func (u UseCase) Read() ([]entities.Actors, error) {
	return u.repo.FindAll()
}

func (u UseCase) ReadBy(col, val string) (*entities.Actors, error) {
	return u.repo.FindBy(col, val)
}

func (u UseCase) Delete(id string) (*entities.Actors, error) {
	return u.repo.SoftDel(id)
}

func (u UseCase) Update(cst *entities.Actors, id string) (*entities.Actors, error) {
	return u.repo.ChangeById(cst, id)
}
func generateHash(data string) string {
	// Membuat objek hash dari algoritma SHA-256
	hash := sha256.New()

	// Mengupdate hash dengan data yang ingin di-hash
	hash.Write([]byte(data))

	// Mengambil nilai hash sebagai array byte
	hashBytes := hash.Sum(nil)

	// Mengubah array byte menjadi representasi heksadesimal
	hashString := fmt.Sprintf("%x", hashBytes)

	return hashString
}
