package model

type Faculty struct {
	ID        uint `json:"id" gorm:"primary_key"`
	Name      string
	Students  []Student `json:"students" gorm:"foreignkey:FacultyID"`
	ManagerID uint
	Manager   Manager `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
