package service

import (
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

	menuMap := make(map[uint]*response.MenuResponse)
	var rootMenus []*response.MenuResponse

	for _, menu := range *menus {
		currentMenu := mapper.MenuToMenuResponse(menu, false)

		if menu.MenuPadreID != nil {
			parent := addParentIfNotExists(menuMap, menu)

			if parent != nil {
				parent.Submenus = append(parent.Submenus, currentMenu)
			}
		}
	}

	for _, menu := range menuMap {
		if menu.MenuPadreId == nil {
			rootMenus = append(rootMenus, menu)
		}
	}

	finalRootMenus := make([]response.MenuResponse, len(rootMenus))
	for i, menu := range rootMenus {
		finalRootMenus[i] = *menu
	}

	return &finalRootMenus, nil
}

func addParentIfNotExists(menuMap map[uint]*response.MenuResponse, menu entity.Menu) *response.MenuResponse {
	if menu.MenuPadreID == nil {
		return nil
	}

	parent, exists := menuMap[*menu.MenuPadreID]
	if exists {
		return parent
	}

	parent = mapper.MenuToMenuResponse(*menu.MenuPadre, true)
	menuMap[*menu.MenuPadreID] = parent

	if menu.MenuPadre.MenuPadreID != nil {
		grandParent := addParentIfNotExists(menuMap, *menu.MenuPadre)
		if grandParent != nil {
			grandParent.Submenus = append(grandParent.Submenus, parent)
		}
	}

	return parent
}
