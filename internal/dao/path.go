package dao

import (
	"ByteScience-WAM-Admin/internal/model/entity"
	"ByteScience-WAM-Admin/pkg/db"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

// PathDao 路径数据访问对象
type PathDao struct{}

// NewPathDao 创建 PathDao 实例
func NewPathDao() *PathDao {
	return &PathDao{}
}

// Insert 插入路径记录
func (pd *PathDao) Insert(ctx context.Context, path *entity.Paths) error {
	return db.Client.WithContext(ctx).Create(path).Error
}

// GetByID 根据 ID 获取路径
func (pd *PathDao) GetByID(ctx context.Context, id string) (*entity.Paths, error) {
	var path entity.Paths
	err := db.Client.WithContext(ctx).
		Where(entity.PathsColumns.ID+" = ?", id).
		Where(entity.PathsColumns.DeletedAt + " IS NULL").
		First(&path).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &path, err
}

// GetByMenuID 根据菜单ID获取路径列表
func (pd *PathDao) GetByMenuID(ctx context.Context, menuID string) ([]*entity.Paths, error) {
	var paths []*entity.Paths
	err := db.Client.WithContext(ctx).
		Where(entity.PathsColumns.MenuID+" = ?", menuID).
		Where(entity.PathsColumns.DeletedAt + " IS NULL").
		Find(&paths).Error
	return paths, err
}

// Update 更新路径信息
func (pd *PathDao) Update(ctx context.Context, id string, updates map[string]interface{}) error {
	return db.Client.WithContext(ctx).
		Model(&entity.Paths{}).
		Where(entity.PathsColumns.ID+" = ?", id).
		Updates(updates).
		Error
}

// SoftDelete 软删除路径记录
func (pd *PathDao) SoftDelete(ctx context.Context, id string) error {
	return db.Client.WithContext(ctx).
		Model(&entity.Paths{}).
		Where(entity.PathsColumns.ID+" = ?", id).
		Update(entity.PathsColumns.DeletedAt, time.Now()).
		Error
}

// Query 分页查询路径
func (pd *PathDao) Query(ctx context.Context, page int, pageSize int, filters map[string]interface{}) ([]*entity.Paths, int64, error) {
	var (
		paths []*entity.Paths
		total int64
	)

	query := db.Client.WithContext(ctx).Model(&entity.Paths{}).Where(entity.PathsColumns.DeletedAt + " IS NULL")

	for key, value := range filters {
		if value != nil && value != "" {
			query = query.Where(key+" = ?", value)
		}
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Scopes(db.PageScope(page, pageSize)).
		Order(entity.PathsColumns.CreatedAt + " DESC").
		Find(&paths).Error; err != nil {
		return nil, 0, err
	}

	return paths, total, nil
}

// GetAll 获取所有路径
func (pd *PathDao) GetAll(ctx context.Context) ([]*entity.Paths, error) {
	var paths []*entity.Paths
	err := db.Client.WithContext(ctx).
		Where(entity.PathsColumns.DeletedAt + " IS NULL").
		Order(entity.PathsColumns.CreatedAt + " DESC").
		Find(&paths).Error
	return paths, err
}