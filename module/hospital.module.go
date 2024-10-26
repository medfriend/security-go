package module

import (
	"gorm.io/gorm"
	"security-go/controller"
	"security-go/repository"
	"security-go/service"
	"security-go/util"
)

func InitializeHospitalModule(db *gorm.DB) *controller.HospitalController {
	return util.GenericModuleInitializer(
		db,
		repository.NewHospitalRepository,
		service.NewHospitalService,
		controller.NewHospitalController,
	)
}
