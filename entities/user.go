package entities

// struktur data suatu table
type User struct {
	ID   uint   `gorm:"primarykey""`
	Name string `json:"name"`
}
