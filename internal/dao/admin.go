package dao

import (
	"ByteScience-WAM-Admin/pkg/db"
	"context"
	"errors"
	"strings"
	"time"

	"ByteScience-WAM-Admin/internal/model/entity"
	"gorm.io/gorm"
)

// InsertAdmin 插入管理员记录
func InsertAdmin(ctx context.Context, admin *entity.Admins) error {
	return db.Client.WithContext(ctx).Create(admin).Error
}

// GetAdminByID 根据 ID 获取管理员
func GetAdminByID(ctx context.Context, id string) (*entity.Admins, error) {
	var admin entity.Admins
	err := db.Client.WithContext(ctx).
		Where(entity.AdminsColumns.ID+" = ?", id).
		Where(entity.AdminsColumns.DeletedAt + " IS NULL").
		First(&admin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &admin, err
}

func GetAdminByFields(ctx context.Context, username, email, phone string) (*entity.Admins, error) {
	// 构建查询条件
	// 基础查询，确保 deleted_at 为 NULL
	query := db.Client.WithContext(ctx).Model(&entity.Admins{}).
		Where(entity.AdminsColumns.DeletedAt + " IS NULL")

	// 创建一个切片来动态构建 OR 条件
	conditions := []string{}
	params := []interface{}{}

	// 动态添加查询条件
	if username != "" {
		conditions = append(conditions, entity.AdminsColumns.Username+" = ?")
		params = append(params, username)
	}
	if phone != "" {
		conditions = append(conditions, entity.AdminsColumns.Phone+" = ?")
		params = append(params, phone)
	}
	if email != "" {
		conditions = append(conditions, entity.AdminsColumns.Email+" = ?")
		params = append(params, email)
	}

	// 如果有条件，使用 OR 连接它们
	if len(conditions) > 0 {
		query = query.Where("("+strings.Join(conditions, " OR ")+")", params...)
	}

	// 执行查询
	var admin entity.Admins
	if err := query.First(&admin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &admin, nil
}

// UpdateAdmin 更新管理员信息
func UpdateAdmin(ctx context.Context, id string, updates map[string]interface{}) error {
	return db.Client.WithContext(ctx).
		Model(&entity.Admins{}).
		Where(entity.AdminsColumns.ID+" = ?", id).
		Updates(updates).
		Error
}

// SoftDeleteAdminByID 软删除管理员记录
func SoftDeleteAdminByID(ctx context.Context, id string) error {
	return db.Client.WithContext(ctx).
		Model(&entity.Admins{}).
		Where(entity.AdminsColumns.ID+" = ?", id).
		Update(entity.AdminsColumns.DeletedAt, time.Now()).
		Error
}

// QueryAdmins 分页查询管理员
func QueryAdmins(ctx context.Context, page int, pageSize int, filters map[string]interface{}) ([]*entity.Admins, int64, error) {
	var (
		admins []*entity.Admins
		total  int64
	)

	query := db.Client.WithContext(ctx).Model(&entity.Admins{}).Where(entity.AdminsColumns.DeletedAt + " IS NULL")

	// 应用过滤条件
	for key, value := range filters {
		if value != nil && value != "" {
			if key == entity.AdminsColumns.Username ||
				key == entity.AdminsColumns.Phone ||
				key == entity.AdminsColumns.Email {
				query = query.Where(key+" LIKE ?", value.(string)+"%")
			} else {
				query = query.Where(key+" = ?", value)
			}
		}
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	if err := query.Scopes(db.PageScope(page, pageSize)).
		Order(entity.AdminsColumns.CreatedAt + " DESC").
		Find(&admins).Error; err != nil {
		return nil, 0, err
	}

	return admins, total, nil
}

// UpdateLastLoginTime 更新管理员的最后登录时间
func UpdateLastLoginTime(ctx context.Context, id string) error {
	return db.Client.WithContext(ctx).
		Model(&entity.Admins{}).
		Where(entity.AdminsColumns.ID+" = ?", id).
		Update(entity.AdminsColumns.LastLoginAt, time.Now()).
		Error
}
