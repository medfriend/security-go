package controller

import (
	"github.com/gin-gonic/gin"
	"security-go/dto"
	"security-go/service"
	"security-go/util"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

// Login realizar la autenticacion del login
// @Summary login
// @Description Autenticacion.
// @Tags auth
// @Accept json
// @Produce json
// @Param entity body dto.AuthDTO true "Información de la entidad"
// @Failure 400 {object} map[string]string "Error en el cuerpo de la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /medfri-security/auth [post]
func (ctrl *AuthController) Login(c *gin.Context) {
	var authDTO dto.AuthDTO

	util.HandlerBindJson(c, &authDTO)

	menus, err := ctrl.AuthService.Auth(&authDTO)

	util.HandlerFoundSuccess(c, err, "auth")
	util.HandlerCreatedSuccess(c, menus)
	
}
