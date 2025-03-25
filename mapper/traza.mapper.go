package mapper

import (
	"security-go/dto"
	"security-go/entity"
)

func MapTrazabilidadToTrazaDTO(traza *entity.TrazabilidadUsuarioAccion) *dto.TrazaDTO {
	return &dto.TrazaDTO{
		Estado:    uint(traza.Estado), // Asegurate de que la conversión es adecuada y segura.
		Coleccion: traza.Coleccion,
	}
}

func MapTrazabilidadToTrazaDTOArray(trazas []entity.TrazabilidadUsuarioAccion) []dto.TrazaDTO {
	dtos := make([]dto.TrazaDTO, len(trazas))
	for i, traza := range trazas {
		dtos[i] = *MapTrazabilidadToTrazaDTO(&traza)
	}
	return dtos
}
