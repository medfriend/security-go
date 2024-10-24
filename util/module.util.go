package util

import "gorm.io/gorm"

type RepositoryConstructor[TRepo any] func(*gorm.DB) TRepo
type ServiceConstructor[TService any, TRepo any] func(TRepo) TService
type ControllerConstructor[TController any, TService any] func(TService) TController

func GenericModuleInitializer[TRepo, TService, TController any](
	db *gorm.DB,
	repoConstructor RepositoryConstructor[TRepo],
	serviceConstructor ServiceConstructor[TService, TRepo],
	controllerConstructor ControllerConstructor[TController, TService],
) TController {

	repository := repoConstructor(db)
	service := serviceConstructor(repository)
	controller := controllerConstructor(service)

	return controller
}
