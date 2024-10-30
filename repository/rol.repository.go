package repository

import (
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/util"
)

type RolRepository interface {
	Save(Rol *entity.Rol) error
	FindById(id uint) (*entity.Rol, error)
	Update(user *entity.Rol) error
	Delete(id uint) error
}

type RolRepositoryImpl struct {
	Base util.BaseRepository[entity.Rol]
}

func NewRolRepository(db *gorm.DB) RolRepository {
	return &RolRepositoryImpl{
		Base: util.BaseRepository[entity.Rol]{DB: db},
	}
}

func (u *RolRepositoryImpl) Save(user *entity.Rol) error {
	return u.Base.Save(user)
}

func (u *RolRepositoryImpl) FindById(id uint) (*entity.Rol, error) {
	return u.Base.FindById(id)
}

func (u *RolRepositoryImpl) Update(Rol *entity.Rol) error {
	return u.Base.Update(Rol)
}

func (u *RolRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}
