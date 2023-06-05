package dto

type CreateRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"  binding:"required"`
	Email     string `json:"email"  binding:"required"`
	Avatar    string `json:"avatar"  binding:"required"`
}
type APIRequest struct {
	Data []CreateRequest `json:"data"`
}
