package repository

import (
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/util"
)

type MenuRepository interface {
	Save(Menu *entity.Menu) error
	FindById(id uint) (*entity.Menu, error)
	FindMenuByResourceAndEntity(resourceIds []uint, entityId uint) (*[]entity.Menu, error)
	GetParentsMenuByEntity(entityId uint) (*[]entity.Menu, error)
	GetChildsMenuByEntity(entityId uint) (*[]entity.Menu, error)
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

func (u *MenuRepositoryImpl) GetParentsMenuByEntity(entityId uint) (*[]entity.Menu, error) {
	var menus []entity.Menu

	err := u.Base.DB.Model(&entity.Menu{}).
		Where("entidad_id = ? AND menu_padre_id IS NULL", entityId).
		Find(&menus).Error

	if err != nil {
		return nil, err
	}

	return &menus, nil
}

func (u *MenuRepositoryImpl) GetChildsMenuByEntity(entityId uint) (*[]entity.Menu, error) {
	var menus []entity.Menu

	err := u.Base.DB.Model(&entity.Menu{}).
		Where("entidad_id = ? AND menu_padre_id IS NOT NULL", entityId).
		Find(&menus).Error

	if err != nil {
		return nil, err
	}

	return &menus, nil
}

func (u *MenuRepositoryImpl) Save(user *entity.Menu) error {
	return u.Base.Save(user)
}

func (u *MenuRepositoryImpl) FindById(id uint) (*entity.Menu, error) {
	return u.Base.FindByIdWithRelations(id, "MenuPadre", "Recurso", "Entidad")
}

func (u *MenuRepositoryImpl) FindMenuByResourceAndEntity(resourceIds []uint, entityId uint) (*[]entity.Menu, error) {
	var menus []entity.Menu

	err := u.Base.DB.Model(&entity.Menu{}).
		Preload("MenuPadre").
		Preload("Recurso").
		Where("recurso_id IN (?) AND entidad_id  = ?", resourceIds, entityId).
		Find(&menus).Error

	if err != nil {
		return nil, err
	}

	return &menus, nil
}

func (u *MenuRepositoryImpl) Update(menu *entity.Menu) error {
	return u.Base.Update(menu)
}

func (u *MenuRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}
