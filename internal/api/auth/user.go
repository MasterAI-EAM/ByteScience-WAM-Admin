package auth

import (
	"ByteScience-WAM-Admin/internal/model/dto"
	"ByteScience-WAM-Admin/internal/model/dto/auth"
	"ByteScience-WAM-Admin/internal/service"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	service *service.UserService
}

// NewUserApi 创建 UserApi 实例并初始化依赖项
func NewUserApi() *UserApi {
	service := service.NewUserService()
	return &UserApi{service: service}
}

// List 获取用户列表
// @Summary 获取用户列表
// @Description 获取用户列表，支持分页，返回用户信息及总条数
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body auth.ListUserRequest true "请求参数"
// @Success 200 {object} auth.ListUserResponse "成功获取用户列表"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误"
// @Router /auth/user [get]
func (api *UserApi) List(ctx *gin.Context, req *auth.ListUserRequest) (res *auth.ListUserResponse, err error) {
	res, err = api.service.List(ctx, req)
	return
}

// Info 获取用户信息
// @Summary 查询用户详情
// @Description 根据用户ID查询用户的详细信息，包括用户的基本信息和角色信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body auth.InfoUserRequest true "查询用户的请求参数"
// @Success 200 {object} auth.InfoUserResponse "成功返回用户详情"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误"
// @Router /auth/user/info [get]
func (api *UserApi) Info(ctx *gin.Context, req *auth.InfoUserRequest) (res *auth.InfoUserResponse, err error) {
	res, err = api.service.Info(ctx, req)
	return
}

// Add 添加用户
// @Summary 添加用户
// @Description 添加一个新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body auth.AddUserRequest true "请求参数"
// @Success 200 {object} dto.Empty "成功添加用户"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误"
// @Router /auth/user [post]
func (api *UserApi) Add(ctx *gin.Context, req *auth.AddUserRequest) (res *dto.Empty, err error) {
	err = api.service.Add(ctx, req)
	return
}

// Edit 编辑用户信息
// @Summary 编辑用户信息
// @Description 编辑指定用户的详细信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body auth.EditUserRequest true "请求参数"
// @Success 200 {object} dto.Empty "成功编辑用户"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误"
// @Router /auth/user [put]
func (api *UserApi) Edit(ctx *gin.Context, req *auth.EditUserRequest) (res *dto.Empty, err error) {
	err = api.service.Edit(ctx, req)
	return
}

// Del 删除用户
// @Summary 删除用户
// @Description 删除指定的用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body auth.DelUserRequest true "请求参数"
// @Success 200 {object} dto.Empty "成功删除用户"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误"
// @Router /auth/user [delete]
func (api *UserApi) Del(ctx *gin.Context, req *auth.DelUserRequest) (res *dto.Empty, err error) {
	err = api.service.Delete(ctx, req)
	return
}

// ResetPassword 重置用户密码
// @Summary 重置用户密码
// @Description 根据提供的用户ID和新密码，重置指定用户的登录密码
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body auth.ResetPasswordRequest true "请求参数，包含要重置密码的用户ID以及新密码"
// @Success 200 {object} dto.Empty "成功重置用户密码，返回空对象表示操作成功"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，如用户ID格式不正确、新密码不符合要求等"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能是数据库更新出错、密码加密失败等情况"
// @Router /auth/user/resetPassword [put]
func (api *UserApi) ResetPassword(ctx *gin.Context, req *auth.ResetPasswordRequest) (res *dto.Empty, err error) {
	err = api.service.ResetPassword(ctx, req)
	return
}
