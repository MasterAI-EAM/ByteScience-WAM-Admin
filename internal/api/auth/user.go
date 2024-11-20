package auth

import (
	"ByteScience-WAM-Admin/internal/model/dto"
	"ByteScience-WAM-Admin/internal/model/dto/auth"
	"ByteScience-WAM-Admin/internal/model/entity"
	"ByteScience-WAM-Admin/pkg/db"
	"ByteScience-WAM-Admin/pkg/redis"
	"github.com/gin-gonic/gin"
)

type UserApi struct{}

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
func (*UserApi) List(ctx *gin.Context, req *auth.ListUserRequest) (res *auth.ListUserResponse, err error) {

	// return nil, utils.NewBusinessError(utils.UserInvalidCredentialsCode)

	redis.Client.Set(ctx, "long", "12312", -1)
	data := make([]entity.Admins, 0)
	db.Client.Scopes(db.PageScope(req.Page, req.PageSize)).Find(&data)
	res = &auth.ListUserResponse{
		Total: 0,
		List:  make([]auth.UserInfo, 0),
	}
	for _, v := range data {
		res.List = append(res.List, auth.UserInfo{
			ID:       v.ID,
			UserName: v.Username,
		})
	}
	// 业务逻辑
	return res, nil
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
func (*UserApi) Info(ctx *gin.Context, req *auth.InfoUserRequest) (res *auth.InfoUserResponse, err error) {
	// 业务逻辑
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
func (*UserApi) Add(ctx *gin.Context, req *auth.AddUserRequest) (res *dto.Empty, err error) {
	// 业务逻辑
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
func (*UserApi) Edit(ctx *gin.Context, req *auth.EditUserRequest) (res *dto.Empty, err error) {
	// 业务逻辑
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
func (*UserApi) Del(ctx *gin.Context, req *auth.DelUserRequest) (res *dto.Empty, err error) {
	// 业务逻辑
	return
}
