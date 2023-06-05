package customers

import "BackendCRM/dto"

type Controller struct {
	useCase *UseCase
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}

func (c Controller) Create(req *dto.CreateRequest) (*dto.CreateResponse, error) {
	user := Customers{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Avatar: req.Avatar}
	err := c.useCase.Create(&user)
	if err != nil {
		return nil, err
	}

	res := &dto.CreateResponse{
		Message: "Success",
		Data: dto.UserItemResponse{
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
	Data []dto.UserItemResponse `json:"data"`
}

func (c Controller) Read() (*ReadResponse, error) {
	users, err := c.useCase.Read()
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, user := range users {
		item := dto.UserItemResponse{
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

func (c Controller) ReadBy(col, val string) (*dto.UserItemResponse, error) {
	user, err := c.useCase.ReadBy(col, val)
	if err != nil {
		return nil, err
	}

	res := &dto.UserItemResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Avatar:    user.Avatar,
	}

	return res, nil
}

func (c Controller) Delete(id string) (*dto.DeleteResponse, error) {
	user, err := c.useCase.Delete(id)
	res := &dto.DeleteResponse{
		Message: "Data berikut berhasil dihapus:",
		Data: dto.UserItemResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}
	return res, err
}

func (c Controller) Update(cst *Customers, id string) (*dto.UpdateResponse, error) {
	user, err := c.useCase.Update(cst, id)
	res := &dto.UpdateResponse{
		Message: "Data berhasil diupdate:",
		Data: dto.UserItemResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}
	return res, err
}
