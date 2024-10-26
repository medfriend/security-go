package repository

import (
	"gorm.io/gorm"
	"security-go/entity"
	"security-go/util"
)

type HospitalRepository interface {
	Save(Hospital *entity.Hospital) error
	FindById(id uint) (*entity.Hospital, error)
	Update(Hospital *entity.Hospital) error
	Delete(id uint) error
}

type HospitalRepositoryImpl struct {
	Base util.BaseRepository[entity.Hospital]
}

func NewHospitalRepository(db *gorm.DB) HospitalRepository {
	return &HospitalRepositoryImpl{
		Base: util.BaseRepository[entity.Hospital]{DB: db},
	}
}

func (u *HospitalRepositoryImpl) Save(Hospital *entity.Hospital) error {
	return u.Base.Save(Hospital)
}

func (u *HospitalRepositoryImpl) FindById(id uint) (*entity.Hospital, error) {
	return u.Base.FindById(id)
}

func (u *HospitalRepositoryImpl) Update(Hospital *entity.Hospital) error {
	return u.Base.Update(Hospital)
}

func (u *HospitalRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}
