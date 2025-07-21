package handlers

import (
	"github.com/gin-gonic/gin"
	"stvCms/internal/services"
)

type LoginAndRegisterHandler struct {
	service services.ILoginAndRegisterService
}

func NewLoginAndRegisterHandler() *LoginAndRegisterHandler {
	return &LoginAndRegisterHandler{
		service: services.NewLoginAndRegisterService(),
	}
}

func (h *LoginAndRegisterHandler) Login(ctx *gin.Context) {

}
