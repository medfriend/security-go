package service

import (
	"security-go/entity"
	"security-go/repository"
)

type HospitalService interface {
	CreateHospital(Hospital *entity.Hospital) error
	GetHospitalById(id uint) (*entity.Hospital, error)
	UpdateHospital(Hospital *entity.Hospital) error
	DeleteHospital(id uint) error
}

type hospitalServiceImpl struct {
	hospitalRepo repository.HospitalRepository
}

func NewHospitalService(hospitalRepo repository.HospitalRepository) HospitalService {
	return &hospitalServiceImpl{
		hospitalRepo: hospitalRepo,
	}
}

func (s *hospitalServiceImpl) CreateHospital(Hospital *entity.Hospital) error {
	return s.hospitalRepo.Save(Hospital)
}

func (s *hospitalServiceImpl) GetHospitalById(id uint) (*entity.Hospital, error) {
	return s.hospitalRepo.FindById(id)
}

func (s *hospitalServiceImpl) UpdateHospital(Hospital *entity.Hospital) error {
	return s.hospitalRepo.Update(Hospital)
}

func (s *hospitalServiceImpl) DeleteHospital(id uint) error {
	return s.hospitalRepo.Delete(id)
}
