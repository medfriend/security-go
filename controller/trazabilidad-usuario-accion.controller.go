package controller

import (
	"github.com/gin-gonic/gin"
	"security-go/service"
	"security-go/util"
)

type TrazabilidadUsuarioAccionController struct {
	trazaService service.TrazabilidadUsuarioAccionService
}

func NewTrazabilidadUsuarioAccionController(trazaService service.TrazabilidadUsuarioAccionService) *TrazabilidadUsuarioAccionController {
	return &TrazabilidadUsuarioAccionController{
		trazaService: trazaService,
	}
}

// GetTrazaByUserId obtiene ultimas traza por usuario_id
// @Summary      Obtener ultimas traza por usuario_id
// @Security      BearerAuth
// @Description  obtiene las ultimas traza por usuarioId.
// @Tags         trazabilidadUsuarioAccion
// @Accept       json
// @Produce      json
// @Param        id  path      uint  true  "ID del usuario"
// @Router       /trazabilidadUsuarioAccion/GetTrazaByUserId/{id} [get]
func (ctrl *TrazabilidadUsuarioAccionController) GetTrazaByUserId(c *gin.Context) {
	id, err := util.StringToUint(c.Param("id"))
	traza, err := ctrl.trazaService.Find(id)
	util.HandlerFoundSuccess(c, err, "traza")
	util.HandlerCreatedSuccess(c, traza, 0)
}
