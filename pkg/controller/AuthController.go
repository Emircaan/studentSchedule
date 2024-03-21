package controller

import (
	"net/http"

	"github.com/emircaan/scheduleManager/pkg/model"
	"github.com/emircaan/scheduleManager/pkg/service"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

func (c *AuthController) Login(ctx echo.Context) error {
	loginRequest := new(model.LoginRequest)
	if err := ctx.Bind(loginRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Geçersiz istek"})
	}
	user, err := c.AuthService.Authentication(loginRequest.Eposta, loginRequest.Sifre)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}
	token, err := c.AuthService.GenerateJWT(user)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "JWT oluşturulamadı"})
	}
	return ctx.JSON(http.StatusOK, map[string]string{"token": token})

}
