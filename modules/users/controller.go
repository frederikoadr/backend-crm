package users

type Controller struct {
	useCase *UseCase
}

func NewController(useCase *UseCase) *Controller {
	return &Controller{
		useCase: useCase,
	}
}

type CreateResponse struct {
	Message string           `json:"message"`
	Data    UserItemResponse `json:"data"`
}

type DeleteResponse struct {
	Message string           `json:"message"`
	Data    UserItemResponse `json:"data"`
}
type UpdateResponse struct {
	Message string           `json:"message"`
	Data    UserItemResponse `json:"data"`
}

type CollectionItemResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
}

type UserItemResponse struct {
	ID          uint                     `json:"id"`
	FirstName   string                   `json:"first_name"`
	LastName    string                   `json:"last_name"`
	Email       string                   `json:"email"`
	Avatar      string                   `json:"avatar"`
	Collections []CollectionItemResponse `json:"collections"`
}

func (c Controller) Create(req *CreateRequest) (*CreateResponse, error) {
	user := Customers{FirstName: req.FirstName, LastName: req.LastName, Email: req.Email, Avatar: req.Avatar}
	err := c.useCase.Create(&user)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{
		Message: "Success",
		Data: UserItemResponse{
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
	Data []UserItemResponse `json:"data"`
}

func (c Controller) Read() (*ReadResponse, error) {
	users, err := c.useCase.Read()
	if err != nil {
		return nil, err
	}

	res := &ReadResponse{}
	for _, user := range users {
		item := UserItemResponse{
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

func (c Controller) Delete(id string) (*DeleteResponse, error) {
	user, err := c.useCase.Delete(id)
	res := &DeleteResponse{
		Message: "Data berikut berhasil dihapus:",
		Data: UserItemResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}
	return res, err
}

func (c Controller) Update(cst *Customers, id string) (*UpdateResponse, error) {
	user, err := c.useCase.Update(cst, id)
	res := &UpdateResponse{
		Message: "Data berhasil diupdate:",
		Data: UserItemResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Avatar:    user.Avatar,
		},
	}
	return res, err
}
