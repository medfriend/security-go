package repository

import (
	"gorm.io/gorm"
	"security-go/dto"
	"security-go/entity"
	"security-go/util"
)

type RolRepository interface {
	Save(Rol *entity.Rol) error
	FindById(id uint) (*entity.Rol, error)
	Find(paginacion dto.PaginationDTO) (dto.PaginatedResponse, error)
	Update(rol *entity.Rol) error
	Delete(id uint) error
}

type RolRepositoryImpl struct {
	Base util.BaseRepository[entity.Rol]
}

func (u *RolRepositoryImpl) Find(paginacion dto.PaginationDTO) (dto.PaginatedResponse, error) {
	return u.Base.Pagination(paginacion)
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
