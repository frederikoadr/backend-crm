package customers

import (
	"BackendCRM/dto"
	"BackendCRM/entities"
)

type Controller struct {
	useCase *UseCase
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}

func (c Controller) Create(req *dto.RequestCustomer) (*dto.CustomerDataResponse, error) {
	user := entities.Customers{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Avatar: req.Avatar}
	err := c.useCase.Create(&user)
	if err != nil {
		return nil, err
	}

	res := &dto.CustomerDataResponse{
		Message: "Success",
		Data: dto.CustomerItemResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}

	return res, nil
}

type ReadResponse struct {
	Data []dto.CustomerItemResponse `json:"data"`
}

func (c Controller) Read() (*ReadResponse, error) {
	users, err := c.useCase.Read()
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, user := range users {
		item := dto.CustomerItemResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		}
		//for _, collection := range user.Collections {
		//	item.Collections = append(item.Collections, CollectionItemResponse{
		//		ID:        collection.ID,
		//		FirstName: collection.FirstName,
		//	})
		//}
		res.Data = append(res.Data, item)
	}

	return res, nil
}

func (c Controller) ReadBy(col, val string) (*dto.CustomerItemResponse, error) {
	user, err := c.useCase.ReadBy(col, val)
	if err != nil {
		return nil, err
	}

	res := &dto.CustomerItemResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
	}

	return res, nil
}

func (c Controller) Delete(id string) (*dto.CustomerDataResponse, error) {
	user, err := c.useCase.Delete(id)
	res := &dto.CustomerDataResponse{
		Message: "Data berikut berhasil dihapus:",
		Data: dto.CustomerItemResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}
	return res, err
}

func (c Controller) Update(req *dto.RequestCustomer, id string) (*dto.CustomerDataResponse, error) {
	cstr := entities.Customers{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Avatar: req.Avatar}
	user, err := c.useCase.Update(&cstr, id)
	res := &dto.CustomerDataResponse{
		Message: "Data berhasil diupdate:",
		Data: dto.CustomerItemResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}
	return res, err
}
