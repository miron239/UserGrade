package userModel

import "gorm.io/gorm"

type UserGrade struct {
	gorm.Model
	UserId        string `gorm:"unique"`
	PostpaidLimit int
	Spp           int
	ShippingFee   int
	ReturnFee     int
}
