package dto

type ErrorResponse struct {
	Error string `json:"error"`
}
type CustomerDataResponse struct {
	Message string               `json:"message"`
	Data    CustomerItemResponse `json:"data"`
}
type ActorDataResponse struct {
	Message string            `json:"message"`
	Data    ActorItemResponse `json:"data"`
}
type CollectionItemResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
}
type CustomerItemResponse struct {
	ID          uint                     `json:"id"`
	FirstName   string                   `json:"first_name"`
	LastName    string                   `json:"last_name"`
	Email       string                   `json:"email"`
	Avatar      string                   `json:"avatar"`
	Collections []CollectionItemResponse `json:"collections"`
}
type ActorItemResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   string `json:"role_id"`
	Verified string `json:"verified"`
	Active   string `json:"active"`
}
