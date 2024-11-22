package v1

import (
	"ByteScience-WAM-Admin/conf"
	"ByteScience-WAM-Admin/internal/api/auth"
	"ByteScience-WAM-Admin/internal/utils"
	"ByteScience-WAM-Admin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAuthRouter(routerGroup *gin.RouterGroup) {
	secret := conf.GlobalConf.Jwt.AccessSecret

	authApi := auth.NewAuthApi()
	{
		utils.RegisterRoute(routerGroup, http.MethodPost, "/login", authApi.Login)
		utils.RegisterRoute(routerGroup, http.MethodPut, "/changPassword", authApi.ChangPassword)
	}

	// 假设您有一个路由组：/auth
	authGroup := routerGroup.Group("/auth", middleware.JWTAuth(secret))
	{
		adminApi := auth.NewAdminApi()
		utils.RegisterRoute(authGroup, http.MethodGet, "/admin", adminApi.List)
		utils.RegisterRoute(authGroup, http.MethodPost, "/admin", adminApi.Add)
		utils.RegisterRoute(authGroup, http.MethodPut, "/admin", adminApi.Edit)
		utils.RegisterRoute(authGroup, http.MethodDelete, "/admin", adminApi.Del)

		userApi := auth.NewUserApi()
		utils.RegisterRoute(authGroup, http.MethodGet, "/user", userApi.List)
		utils.RegisterRoute(authGroup, http.MethodGet, "/user/info", userApi.Info)
		utils.RegisterRoute(authGroup, http.MethodPost, "/user", userApi.Add)
		utils.RegisterRoute(authGroup, http.MethodPut, "/user", userApi.Edit)
		utils.RegisterRoute(authGroup, http.MethodDelete, "/user", userApi.Del)
		utils.RegisterRoute(authGroup, http.MethodPut, "/user/resetPassword", userApi.ResetPassword)

		roleApi := auth.NewRoleApi()
		utils.RegisterRoute(authGroup, http.MethodGet, "/role", roleApi.List)
		utils.RegisterRoute(authGroup, http.MethodGet, "/role/info", roleApi.Info)
		utils.RegisterRoute(authGroup, http.MethodPost, "/role", roleApi.Add)
		utils.RegisterRoute(authGroup, http.MethodPut, "/role", roleApi.Edit)
		utils.RegisterRoute(authGroup, http.MethodDelete, "/role", roleApi.Del)

		menuApi := auth.NewMenuApi()
		utils.RegisterRoute(authGroup, http.MethodGet, "/menu/tree", menuApi.MenuTree)
	}

}
