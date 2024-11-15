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
// @Param request query auth.ListRequest true "请求参数"
// @Success 200 {object} auth.ListResponse "成功获取用户列表"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误"
// @Router /auth/user [get]
func (*UserApi) List(ctx *gin.Context, req *auth.ListRequest) (res *auth.ListResponse, err error) {

	// return nil, utils.NewBusinessError(utils.UserInvalidCredentialsCode)

	redis.Client.Set(ctx, "long", "12312", -1)
	data := make([]entity.Admin, 0)
	db.Client.Scopes(db.PageScope(req.Page, req.PageSize)).Find(&data)
	res = &auth.ListResponse{
		Total: 0,
		Data:  make([]*auth.UserInfo, 0),
	}
	for _, v := range data {
		res.Data = append(res.Data, &auth.UserInfo{
			ID:   v.ID,
			Name: v.Username,
		})
	}
	// 业务逻辑
	return res, nil
}

// Add 添加用户
// @Summary 添加用户
// @Description 添加一个新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body auth.AddRequest true "请求参数"
// @Success 200 {object} dto.EmptyResponse "成功添加用户"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误"
// @Router /auth/user [post]
func (*UserApi) Add(ctx *gin.Context, req *auth.AddRequest) (res *dto.EmptyResponse, err error) {
	// 业务逻辑
	return
}

// Edit 编辑用户信息
// @Summary 编辑用户信息
// @Description 编辑指定用户的详细信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body auth.EditRequest true "请求参数"
// @Success 200 {object} dto.EmptyResponse "成功编辑用户"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误"
// @Router /auth/user [put]
func (*UserApi) Edit(ctx *gin.Context, req *auth.EditRequest) (res *dto.EmptyResponse, err error) {
	// 业务逻辑
	return
}

// Del 删除用户
// @Summary 删除用户
// @Description 删除指定的用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body auth.DelRequest true "请求参数"
// @Success 200 {object} dto.EmptyResponse "成功删除用户"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误"
// @Router /auth/user [delete]
func (*UserApi) Del(ctx *gin.Context, req *auth.DelRequest) (res *dto.EmptyResponse, err error) {
	// 业务逻辑
	return
}
