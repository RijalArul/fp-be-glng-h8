package entity

import (
	"fp-be-glng-h8/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;unique" valid:"required~Your email is required"`
	Email    string `gorm:"not null;unique" valid:"required~Your email is required,email~Invalid email format"`
	Password string `gorm:"not null" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Age      int    `gorm:"not null" valid:"required~Your age is required,range(8|60)~Age minimum Age 8-60"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
