package repository

import (
	"security-go/entity"

	"gorm.io/gorm"
)

type ResourceRepository interface {
	Save(user *entity.Resource) error
	FindById(id uint) (*entity.Resource, error)
	Update(user *entity.Resource) error
	Delete(id uint) error
}

type ResourceRepositoryImpl struct {
	Base BaseRepository[entity.Resource]
}

func NewResourceRepository(db *gorm.DB) ResourceRepository {
	return &ResourceRepositoryImpl{
		Base: BaseRepository[entity.Resource]{DB: db},
	}
}

func (u *ResourceRepositoryImpl) Save(user *entity.Resource) error {
	return u.Base.Save(user)
}

func (u *ResourceRepositoryImpl) FindById(id uint) (*entity.Resource, error) {
	return u.Base.FindById(id)
}

func (u *ResourceRepositoryImpl) Update(user *entity.Resource) error {
	return u.Base.Update(user)
}

func (u *ResourceRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}