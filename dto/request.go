package dto

type RequestCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}
type APIRequest struct {
	Data []RequestCustomer `json:"data"`
}
type RequestActor struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	RoleId   string `json:"role_id"`
	Verified string `json:"verified"`
	Active   string `json:"active"`
}
