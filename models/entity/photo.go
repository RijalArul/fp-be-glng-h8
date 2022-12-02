package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null" valid:"required~Your title is required"`
	Caption  string `gorm:"not null" valid:"required~Your title is required"`
	PhotoUrl string `gorm:"not null" valid:"required~Url of your photo Product is required"`
	UserID   uint
	User     *User
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil

	return
}
