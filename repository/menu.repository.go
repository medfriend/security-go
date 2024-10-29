package repository

import (
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/util"
)

type MenuRepository interface {
	Save(Menu *entity.Menu) error
	FindById(id uint) (*entity.Menu, error)
	Update(user *entity.Menu) error
	Delete(id uint) error
}

type MenuRepositoryImpl struct {
	Base util.BaseRepository[entity.Menu]
}

func NewMenuRepository(db *gorm.DB) MenuRepository {
	return &MenuRepositoryImpl{
		Base: util.BaseRepository[entity.Menu]{DB: db},
	}
}

func (u *MenuRepositoryImpl) Save(user *entity.Menu) error {
	return u.Base.Save(user)
}

func (u *MenuRepositoryImpl) FindById(id uint) (*entity.Menu, error) {
	return u.Base.FindByIdWithRelations(id, "MenuPadre", "Recurso", "Entidad")
}

func (u *MenuRepositoryImpl) Update(user *entity.Menu) error {
	return u.Base.Update(user)
}

func (u *MenuRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}
