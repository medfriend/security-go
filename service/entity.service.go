package service

import (
	"security-go/entity"
	"security-go/repository"
)

type EntityService interface {
	CreateEntity(Entity *entity.Entity) error
	GetEntityById(id uint) (*entity.Entity, error)
	UpdateEntity(Entity *entity.Entity) error
	DeleteEntity(id uint) error
}

type EntityServiceImpl struct {
	EntityRepo repository.EntityRepository
}

func NewEntityService(EntityRepo repository.EntityRepository) EntityService {
	return &EntityServiceImpl{
		EntityRepo: EntityRepo,
	}
}

func (s *EntityServiceImpl) CreateEntity(Entity *entity.Entity) error {
	return s.EntityRepo.Save(Entity)
}

func (s *EntityServiceImpl) GetEntityById(id uint) (*entity.Entity, error) {
	return s.EntityRepo.FindById(id)
}

func (s *EntityServiceImpl) UpdateEntity(Entity *entity.Entity) error {
	return s.EntityRepo.Update(Entity)
}

func (s *EntityServiceImpl) DeleteEntity(id uint) error {
	return s.EntityRepo.Delete(id)
}
