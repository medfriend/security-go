package service

import (
	"fmt"
	"security-go/entity"
	"security-go/mapper"
	"security-go/repository"
	"security-go/response"
)

type MenuService interface {
	CreateMenu(menu *entity.Menu) error
	FindById(id uint) (*entity.Menu, error)
	UpdateMenu(menu *entity.Menu) error
	FindMenuByResourceAndEntity(resourceIds []uint, entityId uint) (*[]response.MenuResponse, error)
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

func (m MenuServiceImpl) FindMenuByResourceAndEntity(resourceIds []uint, entityId uint) (*[]response.MenuResponse, error) {

	menus, _ := m.menuRepository.FindMenuByResourceAndEntity(resourceIds, entityId)

	fmt.Println(menus)

	menuMap := make(map[uint]*response.MenuResponse)

	var rootMenus []*response.MenuResponse

	for _, menu := range *menus {

		currentMenu := mapper.MenuToMenuResponse(menu, false)

		if menu.MenuPadreID != nil {

			parent, exists := menuMap[*menu.MenuPadreID]

			if !exists {

				parent = mapper.MenuToMenuResponse(menu, true)

				menuMap[*menu.MenuPadreID] = parent

				if menu.MenuPadre.MenuPadreID == nil {
					rootMenus = append(rootMenus, parent)
				}
			}

			parent.Submenus = append(parent.Submenus, currentMenu)
		} else {
			rootMenus = append(rootMenus, currentMenu)
		}

		menuMap[menu.MenuID] = currentMenu
	}

	finalRootMenus := make([]response.MenuResponse, len(rootMenus))
	for i, menu := range rootMenus {
		finalRootMenus[i] = *menu
	}

	return &finalRootMenus, nil
}
