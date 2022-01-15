package models

type Staff struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique"`
	Password []byte `json:"-"`
}
