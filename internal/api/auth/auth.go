package auth

import (
	"github.com/gin-gonic/gin"
)

type Api struct{}

func (*Api) Login(ctx *gin.Context) {}

func (*Api) LogOut(ctx *gin.Context) {}
