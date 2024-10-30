package repository

import (
	"security-go/entity"
	"security-go/util"

	"gorm.io/gorm"
)

type ResourcePermissionRepository interface {
	Save(ResourcePermission *entity.ResourcePermission) error
	FindById(id uint) (*entity.ResourcePermission, error)
	Update(user *entity.ResourcePermission) error
	Delete(id uint) error
}

type ResourcePermissionRepositoryImpl struct {
	Base util.BaseRepository[entity.ResourcePermission]
}

func NewResourcePermissionRepository(db *gorm.DB) ResourcePermissionRepository {
	return &ResourcePermissionRepositoryImpl{
		Base: util.BaseRepository[entity.ResourcePermission]{DB: db},
	}
}

func (u *ResourcePermissionRepositoryImpl) Save(user *entity.ResourcePermission) error {
	return u.Base.Save(user)
}

func (u *ResourcePermissionRepositoryImpl) FindById(id uint) (*entity.ResourcePermission, error) {
	return u.Base.FindByIdWithRelations(id, "Recurso", "Permiso")
}

func (u *ResourcePermissionRepositoryImpl) Update(resourcePermission *entity.ResourcePermission) error {
	return u.Base.Update(resourcePermission)
}

func (u *ResourcePermissionRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}