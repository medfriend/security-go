package repository

import (
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/util"
)

type UserRepository interface {
	Save(user *entity.User) error
	FindById(id uint) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
}

type UserRepositoryImpl struct {
	Base util.BaseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Base: util.BaseRepository[entity.User]{DB: db},
	}
}

func (u *UserRepositoryImpl) Save(user *entity.User) error {
	return u.Base.Save(user)
}

func (u *UserRepositoryImpl) FindById(id uint) (*entity.User, error) {
	return u.Base.FindById(id)
}

func (u *UserRepositoryImpl) Update(user *entity.User) error {
	return u.Base.Update(user)
}

func (u *UserRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}
