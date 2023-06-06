package account

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

func (c Controller) Create(req *dto.RequestActor) (*dto.ActorDataResponse, error) {
	user := entities.Actors{Username: req.Username, Password: req.Password, RoleId: req.RoleId, Verified: req.Verified, Active: req.Active}
	err := c.useCase.Create(&user)
	if err != nil {
		return nil, err
	}

	res := &dto.ActorDataResponse{
		Message: "Success",
		Data: dto.ActorItemResponse{
			ID:       user.ID,
			Username: user.Username,
			Password: user.Password,
			RoleId:   user.RoleId,
			Verified: user.Verified,
			Active:   user.Active,
		},
	}

	return res, nil
}

type ReadResponse struct {
	Data []dto.ActorItemResponse `json:"data"`
}

func (c Controller) Read() (*ReadResponse, error) {
	users, err := c.useCase.Read()
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, user := range users {
		item := dto.ActorItemResponse{
			ID:       user.ID,
			Username: user.Username,
			Password: user.Password,
			RoleId:   user.RoleId,
			Verified: user.Verified,
			Active:   user.Active,
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

func (c Controller) ReadBy(col, val string) (*dto.ActorItemResponse, error) {
	user, err := c.useCase.ReadBy(col, val)
	if err != nil {
		return nil, err
	}

	res := &dto.ActorItemResponse{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		RoleId:   user.RoleId,
		Verified: user.Verified,
		Active:   user.Active,
	}

	return res, nil
}

func (c Controller) Delete(id string) (*dto.ActorDataResponse, error) {
	user, err := c.useCase.Delete(id)
	res := &dto.ActorDataResponse{
		Message: "Data berikut berhasil dihapus:",
		Data: dto.ActorItemResponse{
			ID:       user.ID,
			Username: user.Username,
			Password: user.Password,
			RoleId:   user.RoleId,
			Verified: user.Verified,
			Active:   user.Active,
		},
	}
	return res, err
}

func (c Controller) Update(req *dto.RequestActor, id string) (*dto.ActorDataResponse, error) {
	cstr := entities.Actors{Username: req.Username, Password: req.Password, RoleId: req.RoleId, Verified: req.Verified, Active: req.Active}
	user, err := c.useCase.Update(&cstr, id)
	res := &dto.ActorDataResponse{
		Message: "Data berhasil diupdate:",
		Data: dto.ActorItemResponse{
			ID:       user.ID,
			Username: user.Username,
			Password: user.Password,
			RoleId:   user.RoleId,
			Verified: user.Verified,
			Active:   user.Active,
		},
	}
	return res, err
}
