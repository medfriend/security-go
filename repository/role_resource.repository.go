package repository

import (
	"security-go/entity"
	"security-go/util"

	"gorm.io/gorm"
)

type RoleResourceRepository interface {
	Save(RoleResource *entity.RoleResource) error
	FindById(id uint) (*entity.RoleResource, error)
	Update(user *entity.RoleResource) error
	Delete(id uint) error
}

type RoleResourceRepositoryImpl struct {
	Base util.BaseRepository[entity.RoleResource]
}

func NewRoleResourceRepository(db *gorm.DB) RoleResourceRepository {
	return &RoleResourceRepositoryImpl{
		Base: util.BaseRepository[entity.RoleResource]{DB: db},
	}
}

func (u *RoleResourceRepositoryImpl) Save(user *entity.RoleResource) error {
	return u.Base.Save(user)
}

func (u *RoleResourceRepositoryImpl) FindById(id uint) (*entity.RoleResource, error) {
	return u.Base.FindByIdWithRelations(id, "Rol", "Recurso", "Entidad")
}

func (u *RoleResourceRepositoryImpl) Update(roleResource *entity.RoleResource) error {
	return u.Base.Update(roleResource)
}

func (u *RoleResourceRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}