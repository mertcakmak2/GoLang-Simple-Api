package model

type Student struct {
	ID        int `json:"id" gorm:"primary_key"`
	Name      string
	Address   string
	FacultyID uint `json:"-"`
}
