package auth

import (
	"ByteScience-WAM-Admin/internal/model/dto"
	"ByteScience-WAM-Admin/internal/model/dto/auth"
	"ByteScience-WAM-Admin/internal/service"
	"github.com/gin-gonic/gin"
)

type AdminApi struct{}

// List 获取管理员列表
// @Summary 获取管理员列表
// @Description 根据指定条件获取管理员列表信息
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Param req body auth.ListAdminRequest true "请求参数，包含获取管理员列表所需的筛选条件等信息"
// @Success 200 {object} auth.ListAdminResponse "成功返回管理员列表信息"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，例如请求参数格式不正确或缺少必要参数"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能是数据库查询出错、服务端逻辑异常等情况"
// @Router /auth/admin [get]
func (*AdminApi) List(ctx *gin.Context, req *auth.ListAdminRequest) (res *auth.ListAdminResponse, err error) {
	res, err = service.GetAdminList(ctx, req)
	return
}

// Add 添加管理员
// @Summary 添加管理员
// @Description 添加一个新的管理员账户
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Param req body auth.AddAdminRequest true "请求参数，包含要添加的管理员的相关信息，如用户名、密码、权限等"
// @Success 200 {object} dto.Empty "成功添加管理员，返回空对象表示操作成功"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，如信息填写不完整、格式不正确等"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能在数据插入、权限设置等环节出现问题"
// @Router /auth/admin [post]
func (*AdminApi) Add(ctx *gin.Context, req *auth.AddAdminRequest) (res *dto.Empty, err error) {
	err = service.AddAdmin(ctx, req)
	return
}

// Edit 编辑管理员信息
// @Summary 编辑管理员信息
// @Description 根据提供的信息对指定管理员的相关信息进行编辑修改
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Param req body auth.EditAdminRequest true "请求参数，包含要修改的管理员的新信息以及用于定位该管理员的标识信息"
// @Success 200 {object} dto.Empty "成功编辑管理员信息，返回空对象表示操作成功"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，如修改信息不完整、格式不正确或定位标识错误等"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能在数据更新、权限校验等环节出现问题"
// @Router /auth/admin [put]
func (*AdminApi) Edit(ctx *gin.Context, req *auth.EditAdminRequest) (res *dto.Empty, err error) {
	err = service.EditAdmin(ctx, req)
	return
}

// Del 删除管理员
// @Summary 删除管理员
// @Description 根据提供的标识信息删除指定的管理员账户
// @Tags 管理员管理
// @Accept json
// @Produce json
// @Param req body auth.DelAdminRequest true "请求参数，包含用于定位要删除管理员的标识信息"
// @Success 200 {object} dto.Empty "成功删除管理员，返回空对象表示操作成功"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，如定位标识错误、缺少必要参数等"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能在数据删除、权限校验等环节出现问题"
// @Router /auth/admin [delete]
func (*AdminApi) Del(ctx *gin.Context, req *auth.DelAdminRequest) (res *dto.Empty, err error) {
	err = service.DeleteAdmin(ctx, req)
	return
}
