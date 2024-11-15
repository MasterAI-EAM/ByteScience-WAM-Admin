package auth

import (
	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

func (*RoleApi) List(ctx *gin.Context) {}

func (*RoleApi) Add(ctx *gin.Context) {}

func (*RoleApi) Edit(ctx *gin.Context) {}

func (*RoleApi) Del(ctx *gin.Context) {}
