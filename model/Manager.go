package model

type Manager struct {
	ID   uint `json:"id" gorm:"primary_key"`
	Name string
}
