package repository

import (
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/util"
)

type UserRolRepository interface {
	Save(userRol *entity.UserRol) error
	Delete(id uint) error
}

type UserRolRepositoryImpl struct {
	Base util.BaseRepository[entity.UserRol]
}

func NewUserRolRepository(db *gorm.DB) UserRolRepository {
	return &UserRolRepositoryImpl{
		Base: util.BaseRepository[entity.UserRol]{DB: db},
	}
}

func (u *UserRolRepositoryImpl) Save(user *entity.UserRol) error {
	return u.Base.Save(user)
}

func (u *UserRolRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}
