package repository

import (
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/util"
)

type UserRolRepository interface {
	Save(userRol *entity.UserRol) error
	Delete(id uint) error
	FindRolesByUserID(userId uint) ([]uint, error)
}

type UserRolRepositoryImpl struct {
	Base util.BaseRepository[entity.UserRol]
}

func NewUserRolRepository(db *gorm.DB) UserRolRepository {
	return &UserRolRepositoryImpl{
		Base: util.BaseRepository[entity.UserRol]{DB: db},
	}
}

func (u *UserRolRepositoryImpl) Save(user *entity.UserRol) error {
	return u.Base.Save(user)
}

func (u *UserRolRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}

func (u *UserRolRepositoryImpl) FindRolesByUserID(userId uint) ([]uint, error) {

	var rolIDS []uint
	if err := u.Base.DB.Model(&entity.UserRol{}).
		Select("rol_id").
		Where("usuario_id = ?", userId).
		Scan(&rolIDS).Error; err != nil {
		return nil, err // Devuelve nil y el error si ocurre
	}
	return rolIDS, nil
}
