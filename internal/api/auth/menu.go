package auth

import (
	"github.com/gin-gonic/gin"
)

type MenuApi struct{}

func (*MenuApi) List(ctx *gin.Context) {}

func (*MenuApi) Add(ctx *gin.Context) {}

func (*MenuApi) Edit(ctx *gin.Context) {}

func (*MenuApi) Del(ctx *gin.Context) {}
