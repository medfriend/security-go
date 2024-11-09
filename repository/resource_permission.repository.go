package repository

import (
	"security-go/entity"
	"security-go/util"

	"gorm.io/gorm"
)

type ResourcePermissionRepository interface {
	Save(ResourcePermission *entity.ResourcePermission) error
	FindById(id uint) (*entity.ResourcePermission, error)
	FindPermissionByResourceAndRole(resourceIds []uint, roleId []uint) (map[uint][]string, error)
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

func (u *ResourcePermissionRepositoryImpl) FindPermissionByResourceAndRole(resourceIds []uint, rolesIds []uint) (map[uint][]string, error) {
	var permissions []entity.ResourcePermission

	err := u.Base.DB.
		Model(&entity.ResourcePermission{}).
		Preload("Permiso").
		Where("recurso_id IN (?) AND rol_id IN (?)", resourceIds, rolesIds).
		Find(&permissions).Error

	if err != nil {
		return nil, err
	}

	permissionsMap := make(map[uint][]string)

	for _, permission := range permissions {
		permissionsMap[*permission.RecursoID] = append(permissionsMap[*permission.RecursoID], permission.Permiso.Accion)
	}

	return permissionsMap, nil
}

func (u *ResourcePermissionRepositoryImpl) Update(resourcePermission *entity.ResourcePermission) error {
	return u.Base.Update(resourcePermission)
}

func (u *ResourcePermissionRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}
