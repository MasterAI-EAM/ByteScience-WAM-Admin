package auth

import (
	"ByteScience-WAM-Admin/internal/model/dto"
	"ByteScience-WAM-Admin/internal/model/dto/auth"
	"ByteScience-WAM-Admin/internal/service"
	"github.com/gin-gonic/gin"
)

// RoleApi 结构体，保存服务和 DAO 实例
type RoleApi struct {
	service *service.RoleService
}

// NewRoleApi 创建 RoleApi 实例并初始化依赖项
func NewRoleApi() *RoleApi {
	roleService := service.NewRoleService()
	return &RoleApi{service: roleService}
}

// List 获取角色列表
// @Summary 获取角色列表
// @Description 根据指定条件获取系统中的角色列表信息
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param req body auth.ListRoleRequest true "请求参数，包含获取角色列表所需的筛选条件等信息"
// @Success 200 {object} auth.ListRoleResponse "成功返回角色列表信息"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，例如请求参数格式不正确或缺少必要参数"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能是数据库查询出错、服务端逻辑异常等情况"
// @Router /role/list [get]
func (api *RoleApi) List(ctx *gin.Context, req *auth.ListRoleRequest) (res *auth.ListRoleResponse, err error) {
	res, err = api.service.List(ctx, req)
	return
}

// Info 获取角色详细信息
// @Summary 获取角色详细信息
// @Description 根据提供的角色标识获取指定角色的详细信息
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param req body auth.InfoRoleRequest true "请求参数，包含用于定位角色的标识信息等"
// @Success 200 {object} auth.InfoRoleResponse "成功返回指定角色的详细信息"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，例如标识信息错误、格式不正确等"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能是数据库查询出错、服务端逻辑异常等情况"
// @Router /auth/role/info [get]
func (api *RoleApi) Info(ctx *gin.Context, req *auth.InfoRoleRequest) (res *auth.InfoRoleResponse, err error) {
	res, err = api.service.Info(ctx, req)
	return
}

// Add 添加角色
// @Summary 添加角色
// @Description 在系统中添加一个新的角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param req body auth.AddRoleRequest true "请求参数，包含要添加角色的相关信息，如角色名称、角色权限等"
// @Success 200 {object} dto.Empty "成功添加角色，返回空对象表示操作成功"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，如信息填写不完整、格式不正确等"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能在数据插入、权限设置等环节出现问题"
// @Router /auth/role [post]
func (api *RoleApi) Add(ctx *gin.Context, req *auth.AddRoleRequest) (res *dto.Empty, err error) {
	err = api.service.Add(ctx, req)
	return
}

// Edit 编辑角色信息
// @Summary 编辑角色信息
// @Description 根据提供的信息对指定角色的相关信息进行编辑修改
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param req body auth.EditRoleRequest true "请求参数，包含要修改的角色的新信息以及用于定位该角色的标识信息"
// @Success 200 {object} dto.Empty "成功编辑角色信息，返回空对象表示操作成功"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，如修改信息不完整、格式不正确或定位标识错误等"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能在数据更新、权限校验等环节出现问题"
// @Router /auth/role [put]
func (api *RoleApi) Edit(ctx *gin.Context, req *auth.EditRoleRequest) (res *dto.Empty, err error) {
	err = api.service.Edit(ctx, req)
	return
}

// Del 删除角色
// @Summary 删除角色
// @Description 根据提供的标识信息删除指定的角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param req body auth.DelRoleRequest true "请求参数，包含用于定位要删除角色的标识信息"
// @Success 200 {object} dto.Empty "成功删除角色，返回空对象表示操作成功"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，如定位标识错误、缺少必要参数等"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能在数据删除、权限校验等环节出现问题"
// @Router /auth/role [delete]
func (api *RoleApi) Del(ctx *gin.Context, req *auth.DelRoleRequest) (res *dto.Empty, err error) {
	err = api.service.Delete(ctx, req)
	return
}
