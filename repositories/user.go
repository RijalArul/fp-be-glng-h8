package repositories

import (
	"fp-be-glng-h8/models/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(entity.User) (entity.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Create(user entity.User) (entity.User, error) {
	err := r.DB.Create(&user).Error
	return user, err
}
