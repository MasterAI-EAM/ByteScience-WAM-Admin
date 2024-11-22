package service

import (
	"context"

	"ByteScience-WAM-Admin/internal/dao"
	"ByteScience-WAM-Admin/internal/model/dto/auth"
	"ByteScience-WAM-Admin/internal/model/entity"
)

type MenuService struct {
	menuDao *dao.MenuDao
	pathDao *dao.PathDao
}

// NewMenuService 创建一个新的 MenuService 实例
func NewMenuService() *MenuService {
	return &MenuService{
		menuDao: dao.NewMenuDao(),
		pathDao: dao.NewPathDao(),
	}
}

// GetMenuPathTree 获取完整的菜单和路径树
func (ms *MenuService) GetMenuPathTree(ctx context.Context) ([]*auth.MenuNode, error) {
	menus, err := ms.menuDao.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	paths, err := ms.pathDao.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	// 构建菜单和路径树
	return buildMenuPathTree(menus, paths), nil
}

// buildMenuPathTree 构建菜单和路径树
func buildMenuPathTree(menus []*entity.Menus, paths []*entity.Paths) []*auth.MenuNode {
	menuMap := make(map[string]*auth.MenuNode)
	for _, menu := range menus {
		menuMap[menu.ID] = &auth.MenuNode{
			BaseNode: auth.BaseNode{
				ID:       menu.ID,
				ParentID: menu.ParentID,
				Name:     menu.Name,
			},
			MenuData: []*auth.MenuNode{},
			Paths:    []*auth.PathInfo{},
		}
	}

	// 将路径添加到菜单下
	for _, path := range paths {
		if menuNode, exists := menuMap[path.MenuID]; exists {
			menuNode.Paths = append(menuNode.Paths, &auth.PathInfo{
				Path:        path.Path,
				Method:      path.Method,
				Description: path.Description,
			})
		}
	}

	// 构建树形结构
	var tree []*auth.MenuNode
	for _, menu := range menus {
		if menu.ParentID == "" || menu.ParentID == "null" {
			tree = append(tree, menuMap[menu.ID])
		} else if parentMenu, exists := menuMap[menu.ParentID]; exists {
			parentMenu.MenuData = append(parentMenu.MenuData, menuMap[menu.ID])
		}
	}

	return tree
}
