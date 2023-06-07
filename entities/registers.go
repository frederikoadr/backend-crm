package entities

type Registers struct {
	ID           uint   `json:"id"`
	AdminId      uint   `json:"admin_id"`
	SuperAdminId uint   `json:"super_admin_id"`
	Status       string `json:"status"`
}
