package auth

// AddUserRequest 是用于新增用户的请求体结构
type AddUserRequest struct {
	// UserName 用户名，必填，长度限制：3-128字符
	// 必须是唯一的，建议使用字母数字组合，避免特殊字符
	UserName string `json:"userName" validate:"required,min=3,max=128" example:"user1"`

	// Email 邮箱，必填，格式验证
	// 邮箱格式必须合法，例如 "user@example.com"
	Email string `json:"email" validate:"omitempty,email" example:"user@example.com"`

	// Phone 手机号码，必填，符合国际电话号码标准（E.164格式）
	// 例如：+1234567890
	Phone string `json:"phone" validate:"omitempty,e164" example:"+1234567890"`

	// Status 用户状态，必填，1表示启用，0表示禁用
	// 如果未明确设置，则用户默认为启用
	Status int `json:"status" validate:"required,oneof=0 1" example:"1"`

	// Remark 备注，选填，最大长度256字符
	// 用于对用户进行描述或标记
	Remark string `json:"remark" validate:"omitempty,max=256" example:"This is a remark"`

	// RoleIDList 角色ID列表，必填，用于指定该用户具有的角色
	// 用户可以关联多个角色，每个角色的ID必须符合UUID格式
	// 使用 "dive" 校验每个元素是否符合 UUID 格式
	RoleIDList []string `json:"roleIDList" validate:"required,dive,uuid4" example:"role_id_1,role_id_2"`
}

// EditUserRequest 是用于编辑用户信息的请求体结构
type EditUserRequest struct {
	// ID 用户唯一标识，必填，UUID格式
	// 该字段用于指定要编辑的用户，格式必须为UUID
	ID string `json:"id" validate:"required,uuid4" example:"clywh0xv70001rvpgzd6256ns"`

	// Email 邮箱，选填，格式验证
	// 仅在修改邮箱时需要传递，格式必须符合标准邮箱格式
	Email string `json:"email" validate:"omitempty,email" example:"user@example.com"`

	// Phone 手机号码，选填，E.164格式
	// 仅在修改电话号码时需要传递，格式必须符合国际电话规范
	Phone string `json:"phone" validate:"omitempty,e164" example:"+1234567890"`

	// Status 用户状态，必填，1表示启用，0表示禁用
	// 如果没有修改状态，可以使用原值
	Status int `json:"status" validate:"required,oneof=0 1" example:"1"`

	// Remark 备注，选填，最大长度256字符
	// 用于更新用户备注信息
	Remark string `json:"remark" validate:"omitempty,max=256" example:"This is a remark"`

	// RoleIDList 角色ID列表，选填，用于指定该用户具有的角色
	// 用户可以关联多个角色，每个角色的ID必须符合UUID格式
	// 使用 "dive" 校验每个元素是否符合 UUID 格式
	RoleIDList []string `json:"roleIDList" validate:"required,dive,uuid4" example:"role_id_1,role_id_2"`
}

// DelUserRequest 是用于删除用户的请求体结构
type DelUserRequest struct {
	// ID 用户唯一标识，必填，UUID格式
	// 用于指定删除的用户
	ID string `json:"id" validate:"required,uuid4" example:"clywh0xv70001rvpgzd6256ns"`
}

// InfoUserRequest 是用于查询用户详情的请求体结构
type InfoUserRequest struct {
	// ID 用户唯一标识，必填，UUID格式
	// 用于指定删除的用户
	ID string `json:"id" validate:"required,uuid4" example:"clywh0xv70001rvpgzd6256ns"`
}

// ListUserRequest 是用于获取用户列表的请求体结构
type ListUserRequest struct {
	// Page 页码，选填，范围：[1,10000] 默认值1
	// 用于分页查询，最小值为1，最大值为10000
	Page int `json:"page" validate:"omitempty,gte=1,lte=10000" example:"1"`

	// PageSize 页数，选填，范围：[1,10000] 默认值10
	// 用于限制每页显示的用户数量，最小值为1，最大值为10000
	PageSize int `json:"pageSize" validate:"omitempty,gte=1,lte=10000" example:"10"`

	// ID 用户唯一标识，选填，UUID格式
	// 可用于根据用户ID进行过滤查询
	ID string `json:"id" validate:"omitempty,uuid4" example:"clywh0xv70001rvpgzd6256ns"`

	// UserName 用户名，选填，长度限制：3-128字符
	// 用于根据用户名进行过滤查询，最小长度为3，最大长度为128
	UserName string `json:"userName" validate:"omitempty,min=3,max=128" example:"user1"`

	// Email 邮箱，选填，格式验证
	// 用于根据邮箱进行过滤查询，邮箱格式必须合法
	Email string `json:"email" validate:"omitempty,email" example:"user@example.com"`

	// Phone 手机号码，选填，E.164格式
	// 用于根据手机号进行过滤查询，手机号格式应符合国际标准
	Phone string `json:"phone" validate:"omitempty,e164" example:"+1234567890"`

	// Status 用户状态，选填，1表示启用，0表示禁用
	// 用于根据用户状态进行过滤查询
	Status int `json:"status" validate:"omitempty,oneof=0 1" example:"1"`
}

type ListUserResponse struct {
	// total 总条数
	Total int64 `json:"total" example:"100"`
	// List 数据
	List []UserInfo `json:"list"`
}

type UserInfo struct {
	// ID string 编号
	ID string `json:"id" example:"clywh0xv70001rvpgzd6256ns"`
	// UserName string 用户名
	UserName string `json:"userName" example:"user1"`
	// Email string 邮箱
	Email string `json:"email" example:"user@example.com"`
	// Phone string 手机号码
	Phone string `json:"phone" example:"+1234567890"`
	// Status int 状态(1: 启用, 0: 禁用)
	Status int `json:"status" example:"1"`
	// Remark string 备注
	Remark string `json:"remark" example:"This is a remark"`
	// LastLoginAt string 上次登录时间
	LastLoginAt string `json:"lastLoginAt" example:"2024-11-18T15:04:05Z"`
	// CreatedAt string 创建时间
	CreatedAt string `json:"createdAt" example:"2024-11-18T10:00:00Z"`
	// UpdatedAt string 更新时间
	UpdatedAt string `json:"updatedAt" example:"2024-11-18T11:00:00Z"`
}

type InfoUserResponse struct {
	// ID string 编号
	ID string `json:"id" example:"clywh0xv70001rvpgzd6256ns"`
	// UserName string 用户名
	UserName string `json:"userName" example:"user1"`
	// Email string 邮箱
	Email string `json:"email" example:"user@example.com"`
	// Phone string 手机号码
	Phone string `json:"phone" example:"+1234567890"`
	// Status int 状态(1: 启用, 0: 禁用)
	Status int `json:"status" example:"1"`
	// Remark string 备注
	Remark string `json:"remark" example:"This is a remark"`
	// LastLoginAt string 上次登录时间
	LastLoginAt string `json:"lastLoginAt" example:"2024-11-18T15:04:05Z"`
	// CreatedAt string 创建时间
	CreatedAt string `json:"createdAt" example:"2024-11-18T10:00:00Z"`
	// UpdatedAt string 更新时间
	UpdatedAt string `json:"updatedAt" example:"2024-11-18T11:00:00Z"`
	// RoleList 角色列表，包含角色的详细信息
	RoleList []TrimRoleInfo `json:"roleList"`
}

// TrimRoleInfo 用于描述用户角色信息的结构（修剪版）
type TrimRoleInfo struct {
	// ID 角色唯一标识
	ID string `json:"id" example:"role-id-123"`

	// Name 角色名称
	Name string `json:"name" example:"admin"`

	// Description 角色描述
	Description string `json:"description" example:"Administrator role with full access"`
}
