package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/medfriend/shared-commons-go/util/controller"
	"security-go/service"
)

type ParameterController struct {
	ParameterService service.ParameterService
}

func NewParameterController(parameterService service.ParameterService) *ParameterController {
	return &ParameterController{
		ParameterService: parameterService,
	}
}

// FindParameterByCode
// @Summary      Obtener parametro por medio de codigo
// @Security      BearerAuth
// @Description  obtener parametro por medio de codigo
// @Tags         parametros
// @Accept       json
// @Produce      json
// @Param        code  path      uint  true  "ID del menu"
// @Success      200 {object}  dto.FindParameterByCodeDto   "menu encontrado"
// @Router       /parameter/FindParameterByCode/{code} [get]
func (ctrl *ParameterController) FindParameterByCode(c *gin.Context) {
	code := c.Param("code")

	parametros, err := ctrl.ParameterService.FindParameterByCode(code)
	controller.HandlerFoundSuccess(c, err, "parametros")
	controller.HandlerCreatedSuccess(c, parametros, 0)
}
