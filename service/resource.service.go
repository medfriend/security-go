package service

import (
	"security-go/entity"
	"security-go/repository"
)

type ResourceService interface {
	CreateResource(resource *entity.Resource) error
	GetResourceById(id uint) (*entity.Resource, error)
	UpdateResource(resource *entity.Resource) error
	DeleteResource(id uint) error
}

type resourceServiceImpl struct {
	resourceRepo repository.ResourceRepository
}

func NewResourceService(resourceRepo repository.ResourceRepository) ResourceService {
	return &resourceServiceImpl{
		resourceRepo: resourceRepo,
	}
}

func (s *resourceServiceImpl) CreateResource(resource *entity.Resource) error {
	return s.resourceRepo.Save(resource)
}

func (s *resourceServiceImpl) GetResourceById(id uint) (*entity.Resource, error) {
	return s.resourceRepo.FindById(id)
}

func (s *resourceServiceImpl) UpdateResource(resource *entity.Resource) error {
	return s.resourceRepo.Update(resource)
}

func (s *resourceServiceImpl) DeleteResource(id uint) error {
	return s.resourceRepo.Delete(id)
}