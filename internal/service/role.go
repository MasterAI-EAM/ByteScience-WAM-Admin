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

type RoleService struct {
	roleDao           *dao.RoleDao
	menuDao           *dao.MenuDao
	pathDao           *dao.PathDao
	rolePathDao       *dao.RolePathDao
	userRoleDao       *dao.UserRoleDao
	userPermissionDao *dao.UserPermissionDao
}

// NewRoleService 创建一个新的 RoleService 实例
func NewRoleService() *RoleService {
	return &RoleService{
		roleDao:           dao.NewRoleDao(),
		menuDao:           dao.NewMenuDao(),
		pathDao:           dao.NewPathDao(),
		rolePathDao:       dao.NewRolePathDao(),
		userRoleDao:       dao.NewUserRoleDao(),
		userPermissionDao: dao.NewUserPermissionDao(),
	}
}

// Add 添加角色
func (rs *RoleService) Add(ctx context.Context, req *auth.AddRoleRequest) error {
	// 检查是否存在冲突的角色
	conflictingRole, err := addRoleConflictCheck(ctx, req.Name, rs.roleDao)
	if err != nil {
		logger.Logger.Errorf("[AddRole] Error checking role conflict: %v", err)
		return err
	}

	if conflictingRole != nil {
		logger.Logger.Infof("[AddRole] Role name %s already exists", req.Name)
		return utils.NewBusinessError(utils.RoleNameAlreadyExistsCode)
	}

	// 构建角色实体
	role := &entity.Roles{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Description: req.Description,
		Status:      req.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 使用事务闭包
	if err = db.Client.WithContext(ctx).Transaction(func(tx *gorm.DB) error { // 插入角色数据
		if err = rs.roleDao.InsertTx(ctx, tx, role); err != nil {
			logger.Logger.Errorf("[AddRole] Error inserting role into DB: %v", err)
			return err
		}

		// 如果 PathIDList 不为空，则插入角色路径关系
		if len(req.PathIDList) > 0 {
			rolePaths := make([]*entity.RolePaths, 0, len(req.PathIDList))
			for _, pathID := range req.PathIDList {
				rolePaths = append(rolePaths, &entity.RolePaths{
					RoleID: role.ID,
					PathID: pathID,
				})
			}

			if err = rs.rolePathDao.InsertBatchTx(ctx, tx, rolePaths); err != nil {
				logger.Logger.Errorf("[AddRole] Error inserting role paths: %v", err)
				return err
			}
		}

		return nil
	}); err != nil {
		return utils.NewBusinessError(utils.RoleInsertFailedCode)
	}

	return nil
}

// Edit 编辑角色信息
func (rs *RoleService) Edit(ctx context.Context, req *auth.EditRoleRequest) error {
	// 确保角色存在
	role, err := rs.roleDao.GetByID(ctx, req.ID)
	if err != nil {
		logger.Logger.Errorf("[EditRole] Error fetching role by ID: %v", err)
		return utils.NewBusinessError(utils.RoleUpdateFailedCode)
	}
	if role == nil {
		return utils.NewBusinessError(utils.RoleNotFoundCode)
	}

	// 检查是否存在冲突的角色名
	conflictingRole, err := addRoleConflictCheck(ctx, req.Name, rs.roleDao)
	if err != nil {
		logger.Logger.Errorf("[EditRole] Error checking role conflict: %v", err)
		return err
	}
	if conflictingRole != nil && conflictingRole.ID != req.ID {
		logger.Logger.Infof("[EditRole] Role name %s already exists", req.Name)
		return utils.NewBusinessError(utils.RoleNameAlreadyExistsCode)
	}

	// 准备更新字段
	updates := map[string]interface{}{
		entity.RolesColumns.Name:        req.Name,
		entity.RolesColumns.Description: req.Description,
		entity.RolesColumns.UpdatedAt:   time.Now(),
		entity.RolesColumns.Status:      req.Status,
	}

	// 开启事务
	if err = db.Client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 调用 RoleDao 层更新数据
		if err = rs.roleDao.UpdateTx(ctx, tx, req.ID, updates); err != nil {
			logger.Logger.Errorf("[EditRole] Error updating role info in DB: %v", err)
			return err
		}

		// 如果 PathIDList 不为空，则更新角色路径关系
		if req.PathIDList != nil {
			// 删除旧的角色路径关联
			if err = rs.rolePathDao.RemoveByRoleIDTx(ctx, tx, req.ID); err != nil {
				logger.Logger.Errorf("[EditRole] Error deleting old role paths: %v", err)
				return err
			}

			// 插入新的角色路径关联
			rolePaths := make([]*entity.RolePaths, 0, len(req.PathIDList))
			for _, pathID := range req.PathIDList {
				rolePaths = append(rolePaths, &entity.RolePaths{
					RoleID: req.ID,
					PathID: pathID,
				})
			}

			if err = rs.rolePathDao.InsertBatchTx(ctx, tx, rolePaths); err != nil {
				logger.Logger.Errorf("[EditRole] Error inserting new role paths: %v", err)
				return err
			}

			// **移除与该角色关联的用户权限**
			// 找出所有关联该角色的用户
			userIDs, err := rs.userRoleDao.GetUserIDsByRoleIDTx(ctx, tx, req.ID)
			if err != nil {
				logger.Logger.Errorf("[EditRole] Error fetching user IDs for role: %v", err)
				return err
			}

			// 删除受影响用户的权限记录
			if err = rs.userPermissionDao.UpdateUserPermissionsTx(ctx, tx, userIDs); err != nil {
				logger.Logger.Errorf("[EditRole] Error update user permissions: %v", err)
				return err
			}
		}

		return nil
	}); err != nil {
		return utils.NewBusinessError(utils.RoleUpdateFailedCode)
	}

	return nil
}

// Delete 软删除角色
func (rs *RoleService) Delete(ctx context.Context, req *auth.DelRoleRequest) error {
	// 确保角色存在
	role, err := rs.roleDao.GetByID(ctx, req.ID)
	if err != nil {
		logger.Logger.Errorf("[DeleteRole] Error fetching role by ID: %v", err)
		return utils.NewBusinessError(utils.RoleDeleteFailedCode)
	}

	if role == nil {
		return utils.NewBusinessError(utils.RoleNotFoundCode)
	}

	if err := db.Client.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 移除角色路径关联
		if err := rs.rolePathDao.RemoveByRoleIDTx(ctx, tx, req.ID); err != nil {
			logger.Logger.Errorf("[DeleteRole] Error removing role paths: %v", err)
			return err
		}

		// **移除与该角色关联的用户权限**
		// 找出所有关联该角色的用户
		userIDs, err := rs.userRoleDao.GetUserIDsByRoleIDTx(ctx, tx, req.ID)
		if err != nil {
			logger.Logger.Errorf("[DeleteRole] Error fetching user IDs for role: %v", err)
			return err
		}

		// 更新受影响用户的权限记录
		if err = rs.userPermissionDao.UpdateUserPermissionsTx(ctx, tx, userIDs); err != nil {
			logger.Logger.Errorf("[DeleteRole] Error update user permissions: %v", err)
			return err
		}

		// 调用 RoleDao 层进行软删除
		if err = rs.roleDao.SoftDeleteByIDTx(ctx, tx, req.ID); err != nil {
			logger.Logger.Errorf("[DeleteRole] Error soft deleting role: %v", err)
			return err
		}

		return nil
	}); err != nil {
		return utils.NewBusinessError(utils.RoleDeleteFailedCode)
	}

	return nil
}

// Info 角色详情
func (rs *RoleService) Info(ctx context.Context, req *auth.InfoRoleRequest) (*auth.InfoRoleResponse, error) {
	// 根据角色ID获取角色信息
	role, err := rs.roleDao.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, utils.NewBusinessError(utils.RoleNotFoundCode)
	}

	// 获取角色的菜单和路径权限树
	menuPathTree, err := rs.GetRoleMenuPathTree(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// 构造响应数据
	resp := &auth.InfoRoleResponse{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		Status:      role.Status,
		CreatedAt:   role.CreatedAt.Format("2006-01-02T15:04:05Z"),
		UpdatedAt:   role.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		MenuData:    menuPathTree,
	}

	return resp, nil
}

// List 获取角色列表（分页）
func (rs *RoleService) List(ctx context.Context, req *auth.ListRoleRequest) (*auth.ListRoleResponse, error) {
	// 构建过滤条件
	filters := map[string]interface{}{
		entity.RolesColumns.ID:   req.ID,
		entity.RolesColumns.Name: req.Name,
	}

	// 查询数据
	roles, total, err := rs.roleDao.Query(ctx, req.Page, req.PageSize, filters)
	if err != nil {
		logger.Logger.Errorf("[GetRoleList] Error fetching roles: %v", err)
		return nil, utils.NewBusinessError(utils.RoleQueryListFailedCode)
	}

	// 转换数据格式为响应模型
	roleList := make([]auth.RoleInfo, 0)
	for _, role := range roles {
		roleList = append(roleList, auth.RoleInfo{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
			CreatedAt:   role.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   role.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &auth.ListRoleResponse{
		Total: total,
		List:  roleList,
	}, nil
}

// GetRoleMenuPathTree 根据角色ID获取菜单和路径树，并标识权限
func (rs *RoleService) GetRoleMenuPathTree(ctx context.Context, roleID string) ([]*auth.RoleMenuNode, error) {
	menus, err := rs.menuDao.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	paths, err := rs.pathDao.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// 获取角色路径信息
	rolePaths, err := rs.rolePathDao.GetByRoleID(ctx, roleID)
	if err != nil {
		return nil, err
	}

	// 获取路径ID集合
	rolePathIDs := make([]string, len(rolePaths))
	for i, path := range rolePaths {
		rolePathIDs[i] = path.ID
	}

	// 构建角色菜单路径树
	return buildRoleMenuPathTree(menus, paths, rolePathIDs), nil
}

// buildRoleMenuPathTree 构建角色菜单路径树，并标识权限
func buildRoleMenuPathTree(menus []*entity.Menus, paths []*entity.Paths, rolePaths []string) []*auth.RoleMenuNode {
	menuMap := make(map[string]*auth.RoleMenuNode)
	pathPermitted := make(map[string]bool)

	// 初始化菜单节点
	for _, menu := range menus {
		menuMap[menu.ID] = &auth.RoleMenuNode{
			BaseNode: auth.BaseNode{
				ID:       menu.ID,
				ParentID: menu.ParentID,
				Name:     menu.Name,
			},
			IsPermitted: false,
			MenuData:    []*auth.RoleMenuNode{},
			Paths:       []*auth.RolePathInfo{},
		}
	}

	// 添加路径节点并标识权限
	for _, path := range paths {
		isPermitted := utils.Contains(rolePaths, path.ID)
		pathPermitted[path.MenuID] = pathPermitted[path.MenuID] || isPermitted

		if menuNode, exists := menuMap[path.MenuID]; exists {
			menuNode.Paths = append(menuNode.Paths, &auth.RolePathInfo{
				PathInfo: auth.PathInfo{
					Path:        path.Path,
					Method:      path.Method,
					Description: path.Description,
				},
				IsPermitted: isPermitted,
			})
		}
	}

	// 递归向上更新菜单的权限状态
	for _, menu := range menus {
		if menu.ParentID != "" && menu.ParentID != "null" {
			if parentMenu, exists := menuMap[menu.ParentID]; exists {
				parentMenu.IsPermitted = parentMenu.IsPermitted || menuMap[menu.ID].IsPermitted
			}
		}
	}

	// 构建树形结构
	var tree []*auth.RoleMenuNode
	for _, menu := range menus {
		if menu.ParentID == "" || menu.ParentID == "null" {
			tree = append(tree, menuMap[menu.ID])
		} else if parentMenu, exists := menuMap[menu.ParentID]; exists {
			parentMenu.MenuData = append(parentMenu.MenuData, menuMap[menu.ID])
		}
	}

	return tree
}

// addRoleConflictCheck 检查角色名是否冲突
func addRoleConflictCheck(ctx context.Context, roleName string, roleDao *dao.RoleDao) (*entity.Roles, error) {
	conflictingRole, err := roleDao.GetByName(ctx, roleName)
	if err != nil {
		return nil, err
	}
	return conflictingRole, nil
}
