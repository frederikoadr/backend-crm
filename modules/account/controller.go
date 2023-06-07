package account

import (
	"BackendCRM/dto"
	"BackendCRM/entities"
	"strconv"
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

func (c Controller) CreateReg(uid string) (*dto.ActorDataResponse, error) {
	user, err := c.useCase.ActorReadBy("id", uid)
	//user := entities.Actors{Username: req.Username, Password: req.Password, RoleId: req.RoleId, Verified: req.Verified, Active: req.Active}
	reg := entities.Registers{
		AdminId:      user.ID,
		SuperAdminId: 1,
		Status:       "Waiting",
	}
	err = c.useCase.CreateReg(&reg)
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
type ReadRegisResponse struct {
	Data []dto.RegisItemResponse `json:"data"`
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
	user, err := c.useCase.ActorReadBy(col, val)
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
func (c Controller) UpdateReg(id, val string) (*dto.RegisItemResponse, error) {
	uid, err := strconv.ParseUint(id, 0, 64)
	if err != nil {
		return nil, err
	}
	cstr := entities.Registers{AdminId: uint(uid), SuperAdminId: 1, Status: val}
	user, err := c.useCase.UpdateReg(&cstr, id)
	res := &dto.RegisItemResponse{
		ID:           uint(uid),
		AdminId:      user.AdminId,
		SuperAdminId: user.SuperAdminId,
		Status:       user.Status,
	}
	return res, err
}

func (c Controller) ReadRegis() (*ReadRegisResponse, error) {
	users, err := c.useCase.ReadRegis()
	if err != nil {
		return nil, err
	}

	res := &ReadRegisResponse{}
	for _, user := range users {
		item := dto.RegisItemResponse{
			AdminId:      user.AdminId,
			SuperAdminId: user.SuperAdminId,
			Status:       user.Status,
		}
		res.Data = append(res.Data, item)
	}
	return res, nil
}
