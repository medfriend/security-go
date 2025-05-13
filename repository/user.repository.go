package repository

import (
	"gorm.io/gorm"
	"security-go/dto"
	"security-go/entity"
	"security-go/util"
)

type UserRepository interface {
	Save(user *entity.User) error
	FindById(id uint) (*entity.User, error)
	Find(paginacion dto.PaginationDTO, entidad int) (dto.PaginatedResponse, error)
	Query(query string) (*[]entity.User, error)
	FindByUsuario(usuario uint) (*entity.User, error)
	Update(user *dto.UpdateUserDTO) error
	Delete(id uint) error
}

type UserRepositoryImpl struct {
	Base util.BaseRepository[entity.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		Base: util.BaseRepository[entity.User]{DB: db},
	}
}

func (u *UserRepositoryImpl) Query(query string) (*[]entity.User, error) {
	return u.Base.FindAnyField(
		[]string{
			"usuario", "nombre_1", "nombre_2",
			"apellido_paterno", "apellido_materno",
			"email",
		},
		query,
		map[string]bool{
			"usuario": true, // se castea porque es bigint
		})
}

func (u *UserRepositoryImpl) Save(user *entity.User) error {
	return u.Base.Save(user)
}

func (u *UserRepositoryImpl) FindById(id uint) (*entity.User, error) {
	return u.Base.FindById(id)
}

func (u *UserRepositoryImpl) Find(paginacion dto.PaginationDTO, entidad int) (dto.PaginatedResponse, error) {

	if entidad == -1 {
		userPaginacion, err := u.Base.Pagination(paginacion)
		return userPaginacion, err
	}

	userPaginacion, err := u.Base.PaginationWithQuery(paginacion, func(db *gorm.DB) *gorm.DB {
		return db.Table("\"re-usuario-rol\" r").
			Select("r.entidad_id, u.usuario_id, u.usuario, u.nombre_1, u.nombre_2, "+
				"u.apellido_paterno, u.apellido_materno, u.email, u.cambiar_clave, "+
				"u.fecha_cambio_clave, u.fecha_creacion, u.activo, u.fecha_retiro, "+
				"u.tiempo_valides_token, u.clave, u.edad, u.firmado").
			Joins("INNER JOIN usuario u ON u.usuario_id = r.usuario_id").Where("r.entidad_id = ?", entidad)
	})

	return userPaginacion, err
}

func (u *UserRepositoryImpl) FindByUsuario(usuario uint) (*entity.User, error) {
	var usuarioE entity.User

	result := u.Base.DB.Where("usuario = ?", usuario).First(&usuarioE)

	if result.Error != nil {
		return nil, result.Error
	}

	return &usuarioE, nil
}

func (u *UserRepositoryImpl) Update(dto *dto.UpdateUserDTO) error {

	updates := map[string]interface{}{
		"usuario":          dto.Usuario,
		"nombre_1":         dto.Nombre1,
		"nombre_2":         dto.Nombre2,
		"apellido_paterno": dto.ApellidoPaterno,
		"apellido_materno": dto.ApellidoMaterno,
		"email":            dto.Email,
		"edad":             dto.Edad,
		"activo":           dto.Estado,
	}

	result := u.Base.DB.Model(&entity.User{}).Where("usuario_id = ?", dto.Usuario_id).Updates(updates)

	return result.Error
}

func (u *UserRepositoryImpl) Delete(id uint) error {
	return u.Base.Delete(id)
}
