package dto

type ErrorResponse struct {
	Error string `json:"error"`
}
type CreateResponse struct {
	Message string           `json:"message"`
	Data    UserItemResponse `json:"data"`
}
type ReadByResponse struct {
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
