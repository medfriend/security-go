package service

import (
	"security-go/entity"
	"security-go/repository"
)

type MenuService interface {
	CreateMenu(menu *entity.Menu) error
	FindById(id uint) (*entity.Menu, error)
	UpdateMenu(menu *entity.Menu) error
	DeleteMenu(id uint) error
}

type MenuServiceImpl struct {
	menuRepository repository.MenuRepository
}

func NewMenuService(menuRepository repository.MenuRepository) MenuService {
	return &MenuServiceImpl{
		menuRepository: menuRepository,
	}
}

func (m MenuServiceImpl) CreateMenu(menu *entity.Menu) error {
	return m.menuRepository.Save(menu)
}

func (m MenuServiceImpl) FindById(id uint) (*entity.Menu, error) {
	return m.menuRepository.FindById(id)
}

func (m MenuServiceImpl) UpdateMenu(menu *entity.Menu) error {
	return m.menuRepository.Update(menu)
}

func (m MenuServiceImpl) DeleteMenu(id uint) error {
	return m.menuRepository.Delete(id)
}
