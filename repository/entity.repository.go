package repository

import (
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/util"
)

type EntityRepository interface {
	Save(Entity *entity.Entity) error
	FindById(id uint) (*entity.Entity, error)
	Update(Entity *entity.Entity) error
	Delete(id uint) error
}

type EntityRepositoryImpl struct {
	Base util.BaseRepository[entity.Entity]
}

func NewEntityRepository(db *gorm.DB) EntityRepository {
	return &EntityRepositoryImpl{
		Base: util.BaseRepository[entity.Entity]{DB: db},
	}
}

func (u *EntityRepositoryImpl) Save(Entity *entity.Entity) error {
	return u.Base.Save(Entity)
}

func (u *EntityRepositoryImpl) FindById(id uint) (*entity.Entity, error) {
	return u.Base.FindById(id)
}

func (u *EntityRepositoryImpl) Update(Entity *entity.Entity) error {
	return u.Base.Update(Entity)
}

func (u *EntityRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}
