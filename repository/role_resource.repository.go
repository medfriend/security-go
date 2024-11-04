package repository

import (
	"security-go/entity"
	"security-go/util"

	"gorm.io/gorm"
)

type RoleResourceRepository interface {
	Save(RoleResource *entity.RoleResource) error
	FindById(id uint) (*entity.RoleResource, error)
	FindResourceByRoleIds(roleIds []uint) ([]uint, error)
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

func (u *RoleResourceRepositoryImpl) FindResourceByRoleIds(roleIds []uint) ([]uint, error) {
	var resourceIds []uint

	err := u.Base.DB.
		Model(&entity.RoleResource{}).
		Where("rol_id IN (?)", roleIds).
		Pluck("recurso_id", &resourceIds).Error

	if err != nil {
		return nil, err
	}

	return resourceIds, nil
}

func (u *RoleResourceRepositoryImpl) Update(roleResource *entity.RoleResource) error {
	return u.Base.Update(roleResource)
}

func (u *RoleResourceRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}
