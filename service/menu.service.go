package service

import (
	"security-go/entity"
	"security-go/repository"
)

type MenuService interface {
	CreateMenu(menu *entity.Menu) error
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
