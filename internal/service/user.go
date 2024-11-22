package service

import (
	"ByteScience-WAM-Admin/internal/dao"
	"ByteScience-WAM-Admin/internal/model/dto/auth"
	"ByteScience-WAM-Admin/internal/model/entity"
	"ByteScience-WAM-Admin/internal/utils"
	"ByteScience-WAM-Admin/pkg/db"
	"ByteScience-WAM-Admin/pkg/logger"
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
	dao         *dao.UserDao
	userRoleDao *dao.UserRoleDao
	roleDao     *dao.RoleDao
}

// NewUserService 创建一个新的 UserService 实例
func NewUserService() *UserService {
	return &UserService{
		dao:         dao.NewUserDao(),
		userRoleDao: dao.NewUserRoleDao(),
		roleDao:     dao.NewRoleDao(),
	}
}

// 检查用户是否存在的公共方法
func (us *UserService) checkUserExistence(ctx context.Context, userID string) (*entity.Users, error) {
	existingUser, err := us.dao.GetByID(ctx, userID)
	if err != nil {
		logger.Logger.Errorf("[CheckUserExistence] Error retrieving user: %v", err)
		return nil, utils.NewBusinessError(utils.UserQueryFailedCode)
	}
	if existingUser == nil {
		return nil, utils.NewBusinessError(utils.UserNotFoundCode)
	}
	return existingUser, nil
}

// Add 添加用户
func (us *UserService) Add(ctx context.Context, req *auth.AddUserRequest) error {
	// 检查是否存在冲突的记录
	conflictingUser, err := us.dao.GetByFields(ctx, req.UserName, req.Email, req.Phone)
	if err != nil {
		logger.Logger.Errorf("[AddUser] Error checking user conflicts: %v", err)
		return utils.NewBusinessError(utils.UserConflictCheckFailedCode)
	}

	if conflictingUser != nil {
		if conflictingUser.Username == req.UserName {
			logger.Logger.Infof("[AddUser] Username %s already exists", req.UserName)
			return utils.NewBusinessError(utils.UsernameAlreadyExistsCode)
		}
		if conflictingUser.Email == req.Email {
			logger.Logger.Infof("[AddUser] Email %s already exists", req.Email)
			return utils.NewBusinessError(utils.EmailAlreadyExistsCode)
		}
		if conflictingUser.Phone == req.Phone {
			logger.Logger.Infof("[AddUser] Phone %s already exists", req.Phone)
			return utils.NewBusinessError(utils.PhoneAlreadyExistsCode)
		}
	}

	// 构建用户实体
	user := &entity.Users{
		ID:        uuid.New().String(),
		Username:  req.UserName,
		Nickname:  req.Nickname,
		Email:     req.Email,
		Phone:     req.Phone,
		Status:    req.Status,
		Remark:    req.Remark,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// 开启事务
	if err = db.Client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 插入用户
		if err = us.dao.InsertTx(ctx, tx, user); err != nil {
			logger.Logger.Errorf("[AddUser] Error inserting user: %v", err)
			return err
		}

		if len(req.RoleIDList) > 0 {
			// 构建用户角色关联
			var userRoles []*entity.UserRoles
			for _, roleID := range req.RoleIDList {
				userRole := &entity.UserRoles{
					UserID: user.ID,
					RoleID: roleID,
				}
				userRoles = append(userRoles, userRole)
			}

			// 批量插入用户角色关联
			if err = us.userRoleDao.InsertBatchTx(ctx, tx, userRoles); err != nil {
				logger.Logger.Errorf("[AddUser] Error assigning roles: %v", err)
				return err
			}
		}

		return nil
	}); err != nil {
		return utils.NewBusinessError(utils.UserInsertFailedCode)
	}

	return nil
}

// Edit 编辑用户
func (us *UserService) Edit(ctx context.Context, req *auth.EditUserRequest) error {
	// 检查用户是否存在
	_, err := us.checkUserExistence(ctx, req.ID)
	if err != nil {
		return err
	}

	// 检查是否存在冲突的记录
	conflictingUser, err := us.dao.GetByFields(ctx, req.UserName, req.Email, req.Phone)
	if err != nil {
		logger.Logger.Errorf("[EditUser] Error checking user conflicts: %v", err)
		return utils.NewBusinessError(utils.UserConflictCheckFailedCode)
	}
	if conflictingUser != nil && conflictingUser.ID != req.ID {
		if conflictingUser.Username == req.UserName {
			logger.Logger.Infof("[EditUser] Username %s already exists", req.UserName)
			return utils.NewBusinessError(utils.UsernameAlreadyExistsCode)
		}
		if conflictingUser.Email == req.Email {
			logger.Logger.Infof("[EditUser] Email %s already exists", req.Email)
			return utils.NewBusinessError(utils.EmailAlreadyExistsCode)
		}
		if conflictingUser.Phone == req.Phone {
			logger.Logger.Infof("[EditUser] Phone %s already exists", req.Phone)
			return utils.NewBusinessError(utils.PhoneAlreadyExistsCode)
		}
	}

	// 开启事务
	if err = db.Client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 更新用户信息
		updates := map[string]interface{}{
			entity.UsersColumns.Username:  req.UserName,
			entity.UsersColumns.Nickname:  req.Nickname,
			entity.UsersColumns.Email:     req.Email,
			entity.UsersColumns.Phone:     req.Phone,
			entity.UsersColumns.Status:    req.Status,
			entity.UsersColumns.Remark:    req.Remark,
			entity.UsersColumns.UpdatedAt: time.Now(),
		}

		if err = us.dao.UpdateTx(ctx, tx, req.ID, updates); err != nil {
			logger.Logger.Errorf("[EditUser] Error updating user: %v", err)
			return err
		}

		if len(req.RoleIDList) > 0 {
			// 移除旧的角色关联
			if err = us.userRoleDao.RemoveByUserIDTx(ctx, tx, req.ID); err != nil {
				logger.Logger.Errorf("[EditUser] Error removing user roles: %v", err)
				return err
			}

			// 构建新的用户角色关联
			var userRoles []*entity.UserRoles
			for _, roleID := range req.RoleIDList {
				userRole := &entity.UserRoles{
					UserID: req.ID,
					RoleID: roleID,
				}
				userRoles = append(userRoles, userRole)
			}

			// 批量插入新的角色关联
			if err = us.userRoleDao.InsertBatchTx(ctx, tx, userRoles); err != nil {
				logger.Logger.Errorf("[EditUser] Error assigning new roles: %v", err)
				return err
			}
		}

		return nil
	}); err != nil {
		utils.NewBusinessError(utils.UserUpdateFailedCode)
	}

	return nil
}

// Delete 删除用户
func (us *UserService) Delete(ctx context.Context, req *auth.DelUserRequest) error {
	// 检查用户是否存在
	if _, err := us.checkUserExistence(ctx, req.ID); err != nil {
		return err
	}

	if err := db.Client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 执行软删除
		if err := us.dao.SoftDeleteByIDTx(ctx, tx, req.ID); err != nil {
			logger.Logger.Errorf("[DeleteUser] Error deleting user: %v", err)
			return err
		}

		// 移除用户角色关联
		if err := us.userRoleDao.RemoveByUserIDTx(ctx, tx, req.ID); err != nil {
			logger.Logger.Errorf("[DeleteUser] Error removing user roles: %v", err)
			return err
		}

		return nil
	}); err != nil {
		return utils.NewBusinessError(utils.UserDeleteFailedCode)
	}

	return nil
}

// Info 获取用户详细信息
func (us *UserService) Info(ctx context.Context, req *auth.InfoUserRequest) (*auth.InfoUserResponse, error) {
	user, err := us.checkUserExistence(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// 查询用户角色信息
	roles, err := us.userRoleDao.GetRolesByUserID(ctx, req.ID)
	if err != nil {
		logger.Logger.Errorf("[InfoUser] Error retrieving user roles: %v", err)
		return nil, utils.NewBusinessError(utils.UserQueryFailedCode)
	}

	roleList := make([]auth.TrimRoleInfo, len(roles))
	for i, role := range roles {
		roleList[i] = auth.TrimRoleInfo{
			ID:   role.ID,
			Name: role.Name,
		}
	}

	return &auth.InfoUserResponse{
		ID:          user.ID,
		UserName:    user.Username,
		Nickname:    user.Nickname,
		Email:       user.Email,
		Phone:       user.Phone,
		Status:      user.Status,
		Remark:      user.Remark,
		LastLoginAt: user.LastLoginAt.Format(time.RFC3339),
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
		RoleList:    roleList,
	}, nil
}

// List 用户列表
func (us *UserService) List(ctx context.Context, req *auth.ListUserRequest) (*auth.ListUserResponse, error) {
	filters := map[string]interface{}{
		entity.UsersColumns.ID:       req.ID,
		entity.UsersColumns.Username: req.UserName,
		entity.UsersColumns.Email:    req.Email,
		entity.UsersColumns.Phone:    req.Phone,
	}

	if req.Status != nil {
		filters[entity.UsersColumns.Status] = req.Status
	}

	users, total, err := us.dao.Query(ctx, req.Page, req.PageSize, filters)
	if err != nil {
		logger.Logger.Errorf("[ListUser] Error querying users: %v", err)
		return nil, utils.NewBusinessError(utils.UserQueryFailedCode)
	}

	// 构造返回
	userList := make([]auth.UserInfo, 0)
	for _, user := range users {
		userList = append(userList, auth.UserInfo{
			ID:          user.ID,
			UserName:    user.Username,
			Nickname:    user.Nickname,
			Remark:      user.Remark,
			Email:       user.Email,
			Phone:       user.Phone,
			Status:      user.Status,
			CreatedAt:   user.CreatedAt.Format(time.RFC3339),
			LastLoginAt: user.LastLoginAt.Format(time.RFC3339),
		})
	}

	return &auth.ListUserResponse{
		Total: total,
		List:  userList,
	}, nil
}
