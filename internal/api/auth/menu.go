package auth

import (
	"ByteScience-WAM-Admin/internal/model/dto"
	"ByteScience-WAM-Admin/internal/model/dto/auth"
	"github.com/gin-gonic/gin"
)

type MenuApi struct{}

// MenuTree 获取菜单树结构
// @Summary 获取菜单树结构
// @Description 获取系统中以树形结构展示的菜单信息，可用于前端展示菜单层级关系等用途。
// @Tags 菜单管理
// @Accept json
// @Produce json
// @Param _ body dto.Empty true "此参数为空对象，当前操作无需额外传入请求参数"
// @Success 200 {object} auth.MenuTreeResponse "成功返回菜单树结构信息"
// @Failure 400 {object} dto.ErrorResponse "请求参数错误，虽然当前操作此情况一般不会出现，但若传入不符合预期的参数（如非空对象等）则会触发"
// @Failure 500 {object} dto.ErrorResponse "服务器内部错误，可能是数据库查询出错、服务端逻辑异常等情况导致无法正确生成菜单树结构"
// @Router /auth/menu/tree [get]
func (*MenuApi) MenuTree(ctx *gin.Context, _ *dto.Empty) (res *auth.MenuTreeResponse, err error) {
	return
}
