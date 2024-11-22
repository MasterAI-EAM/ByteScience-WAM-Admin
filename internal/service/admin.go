package service

import (
	"ByteScience-WAM-Admin/internal/dao"
	"ByteScience-WAM-Admin/internal/model/dto/auth"
	"ByteScience-WAM-Admin/internal/model/entity"
	"ByteScience-WAM-Admin/internal/utils"
	"ByteScience-WAM-Admin/pkg/logger"
	"context"
	"time"

	"github.com/google/uuid"
)

type AdminService struct {
	dao *dao.AdminDao // 添加 AdminDao 作为成员
}

// NewAdminService 创建一个新的 AdminService 实例
func NewAdminService() *AdminService {
	return &AdminService{
		dao: dao.NewAdminDao(),
	}
}

// Add 添加管理员
func (as *AdminService) Add(ctx context.Context, req *auth.AddAdminRequest) error {
	// 密码加密
	hashedPassword, err := utils.EncryptPassword(req.Password)
	if err != nil {
		// 记录加密错误的详细信息
		logger.Logger.Errorf("[AddAdmin] utils.EncryptPassword error: %v", err)
		return utils.NewBusinessError(utils.PasswordGenerationFailedCode)
	}

	// 检查是否存在冲突的记录
	conflictingAdmin, err := as.dao.GetByFields(ctx, req.UserName, req.Email, req.Phone)
	if err != nil {
		logger.Logger.Errorf("[AddAdmin] Error checking admin conflicts: %v", err)
		return err
	}

	// 根据冲突的字段返回相应的错误
	if conflictingAdmin != nil {
		if conflictingAdmin.Username == req.UserName {
			logger.Logger.Infof("[AddAdmin] Admin username %s already exists", req.UserName)
			return utils.NewBusinessError(utils.AdminUsernameAlreadyExistsCode)
		}
		if conflictingAdmin.Email == req.Email {
			logger.Logger.Infof("[AddAdmin] Admin email %s already exists", req.Email)
			return utils.NewBusinessError(utils.AdminEmailAlreadyExistsCode)
		}
		if conflictingAdmin.Phone == req.Phone {
			logger.Logger.Infof("[AddAdmin] Admin phone %s already exists", req.Phone)
			return utils.NewBusinessError(utils.AdminPhoneAlreadyExistsCode)
		}
	}

	// 构建管理员实体
	admin := &entity.Admins{
		ID:        uuid.New().String(),
		Username:  req.UserName,
		Nickname:  req.Nickname,
		Email:     req.Email,
		Phone:     req.Phone,
		Remark:    req.Remark,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 调用 DAO 层插入数据
	if err = as.dao.Insert(ctx, admin); err != nil {
		// 记录插入数据库的错误
		logger.Logger.Errorf("[AddAdmin] Error inserting admin into DB: %v", err)
		return utils.NewBusinessError(utils.AdminInsertFailedCode) // 新增错误码 AdminInsertFailedCode
	}

	return nil
}

// Edit 编辑管理员信息
func (as *AdminService) Edit(ctx context.Context, req *auth.EditAdminRequest) error {
	// 确保管理员存在
	admin, err := as.dao.GetByID(ctx, req.ID)
	if err != nil {
		logger.Logger.Errorf("[EditAdmin] Error fetching admin by ID: %v", err)
		return err
	}
	if admin == nil {
		return utils.NewBusinessError(utils.AdminNotFoundCode)
	}

	// 检查是否存在冲突的记录
	conflictingAdmin, err := as.dao.GetByFields(ctx, req.UserName, req.Email, req.Phone)
	if err != nil {
		logger.Logger.Errorf("[EditAdmin] Error checking admin conflicts: %v", err)
		return err
	}

	// 如果有冲突的管理员，且不是当前管理员，返回相应的错误
	if conflictingAdmin != nil && conflictingAdmin.ID != req.ID {
		if conflictingAdmin.Username == req.UserName {
			logger.Logger.Infof("[EditAdmin] Admin username %s already exists", req.UserName)
			return utils.NewBusinessError(utils.AdminUsernameAlreadyExistsCode)
		}
		if conflictingAdmin.Email == req.Email {
			logger.Logger.Infof("[EditAdmin] Admin email %s already exists", req.Email)
			return utils.NewBusinessError(utils.AdminEmailAlreadyExistsCode)
		}
		if conflictingAdmin.Phone == req.Phone {
			logger.Logger.Infof("[EditAdmin] Admin phone %s already exists", req.Phone)
			return utils.NewBusinessError(utils.AdminPhoneAlreadyExistsCode)
		}
	}

	// 准备更新字段
	updates := map[string]interface{}{
		entity.AdminsColumns.Username:  req.UserName,
		entity.AdminsColumns.Email:     req.Email,
		entity.AdminsColumns.Phone:     req.Phone,
		entity.AdminsColumns.Nickname:  req.Nickname,
		entity.AdminsColumns.Remark:    req.Remark,
		entity.AdminsColumns.UpdatedAt: time.Now(),
	}

	// 调用 DAO 层更新数据
	if err = as.dao.Update(ctx, req.ID, updates); err != nil {
		// 记录更新管理员信息时的错误
		logger.Logger.Errorf("[EditAdmin] Error updating admin info in DB: %v", err)
		return utils.NewBusinessError(utils.AdminUpdateFailedCode)
	}

	return nil
}

// Delete 软删除管理员
func (as *AdminService) Delete(ctx context.Context, req *auth.DelAdminRequest) error {
	// 确保管理员存在
	admin, err := as.dao.GetByID(ctx, req.ID)
	if err != nil {
		logger.Logger.Errorf("[DeleteAdmin] Error fetching admin by ID: %v", err)
		return err
	}
	if admin == nil {
		return utils.NewBusinessError(utils.AdminNotFoundCode)
	}

	// 调用 DAO 层进行软删除
	if err = as.dao.SoftDeleteByID(ctx, req.ID); err != nil {
		// 记录软删除操作错误
		logger.Logger.Errorf("[DeleteAdmin] Error soft deleting admin: %v", err)
		return utils.NewBusinessError(utils.AdminDeleteFailedCode)
	}

	return nil
}

// GetList 获取管理员列表（分页）
func (as *AdminService) GetList(ctx context.Context, req *auth.ListAdminRequest) (*auth.ListAdminResponse, error) {
	// 构建过滤条件
	filters := map[string]interface{}{
		entity.AdminsColumns.ID:       req.ID,
		entity.AdminsColumns.Username: req.UserName,
		entity.AdminsColumns.Email:    req.Email,
		entity.AdminsColumns.Phone:    req.Phone,
	}

	// 查询数据
	admins, total, err := as.dao.Query(ctx, req.Page, req.PageSize, filters)
	if err != nil {
		// 查询失败时返回具体的业务错误
		logger.Logger.Errorf("[GetAdminList] Error fetching admins: %v", err)
		return nil, utils.NewBusinessError(utils.AdminQueryListFailedCode)
	}

	// 转换数据格式为响应模型
	adminList := make([]auth.AdminInfo, 0)
	for _, admin := range admins {
		adminList = append(adminList, auth.AdminInfo{
			ID:          admin.ID,
			UserName:    admin.Username,
			Nickname:    admin.Nickname,
			Email:       admin.Email,
			Phone:       admin.Phone,
			Remark:      admin.Remark,
			LastLoginAt: admin.LastLoginAt.Format(time.RFC3339),
			CreatedAt:   admin.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   admin.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &auth.ListAdminResponse{
		Total: total,
		List:  adminList,
	}, nil
}

// UpdateLastLoginTime 更新管理员的最后登录时间
func (as *AdminService) UpdateLastLoginTime(ctx context.Context, id string) error {
	// 确保管理员存在
	admin, err := as.dao.GetByID(ctx, id)
	if err != nil {
		return err
	}
	if admin == nil {
		return utils.NewBusinessError(utils.AdminNotFoundCode)
	}

	// 调用 DAO 层更新最后登录时间
	return as.dao.UpdateLastLoginTime(ctx, id)
}
