package entities

type Register struct {
	AdminId      uint   `json:"admin_id"`
	SuperAdminId uint   `json:"super_admin_id"`
	Status       string `json:"status"`
}
