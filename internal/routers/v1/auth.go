package v1

import (
	"ByteScience-WAM-Admin/internal/api/auth"
	"ByteScience-WAM-Admin/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(routerGroup *gin.RouterGroup) {
	userApi := &auth.UserApi{}
	{
		utils.RegisterRoute(routerGroup, http.MethodGet, "/auth/user", userApi.List)
		utils.RegisterRoute(routerGroup, http.MethodPost, "/auth/user", userApi.Add)
		utils.RegisterRoute(routerGroup, http.MethodPut, "/auth/user", userApi.Edit)
		utils.RegisterRoute(routerGroup, http.MethodDelete, "/auth/user", userApi.Del)
	}

	roleApi := &auth.RoleApi{}
	{
		routerGroup.GET("/auth/role", roleApi.List)
		routerGroup.POST("/auth/role", roleApi.Add)
		routerGroup.PUT("/auth/role", roleApi.Edit)
		routerGroup.DELETE("/auth/role", roleApi.Del)
	}

	menuApi := &auth.MenuApi{}
	{
		routerGroup.GET("/auth/menu", menuApi.List)
		routerGroup.POST("/auth/menu", menuApi.Add)
		routerGroup.PUT("/auth/menu", menuApi.Edit)
		routerGroup.DELETE("/auth/menu", menuApi.Del)
	}

	authApi := &auth.Api{}
	{
		routerGroup.POST("/auth/login", authApi.Login)
		routerGroup.POST("/auth/logout", authApi.LogOut)
	}
}
