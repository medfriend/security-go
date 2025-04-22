package repository

import (
	"github.com/medfriend/shared-commons-go/util/repository"
	"gorm.io/gorm"
	"security-go/dto"
	"security-go/entity"
)

type ParameterRepository interface {
	FindParameterByCode(code string) (dto.FindParameterByCodeDto, error)
}
type ParameterRepositoryImpl struct {
	Base repository.BaseRepository[entity.Parameter]
}

func (p ParameterRepositoryImpl) FindParameterByCode(code string) (dto.FindParameterByCodeDto, error) {
	var result dto.FindParameterByCodeDto

	err := p.Base.DB.Table("parametro p").
		Joins("INNER JOIN parametro_detalle pd ON pd.parametro_id = p.parametro_id").
		Select("p.nombre, pd.valor_codigo, pd.valor_descripcion, pd.parametro_detalle_padre_id").
		Where("p.codigo = ?", code).
		Scan(&result).Error

	return result, err
}

func NewParameterRepository(db *gorm.DB) ParameterRepository {
	return &ParameterRepositoryImpl{
		Base: repository.BaseRepository[entity.Parameter]{DB: db},
	}
}
