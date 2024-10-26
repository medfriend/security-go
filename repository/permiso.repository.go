package repository

import (
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/util"
)

type PermisoRepository interface {
	Save(permiso *entity.Permiso) error
	FindById(id uint) (*entity.Permiso, error)
	Update(permiso *entity.Permiso) error
	Delete(id uint) error
}

type PermisoRepositoryImpl struct {
	Base util.BaseRepository[entity.Permiso]
}

func NewPermisoRepository(db *gorm.DB) PermisoRepository {
	return &PermisoRepositoryImpl{
		Base: util.BaseRepository[entity.Permiso]{DB: db},
	}
}

func (r *PermisoRepositoryImpl) Save(permiso *entity.Permiso) error {
	return r.Base.Save(permiso)
}

func (r *PermisoRepositoryImpl) FindById(id uint) (*entity.Permiso, error) {
	return r.Base.FindById(id)
}

func (r *PermisoRepositoryImpl) Update(permiso *entity.Permiso) error {
	return r.Base.Update(permiso)
}

func (r *PermisoRepositoryImpl) Delete(id uint) error {
	return r.Base.Delete(id)
}
