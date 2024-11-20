package auth

// ListRoleRequest 用于查询角色列表的请求体结构
type ListRoleRequest struct {
	// Page 页码，选填，范围限制：[1,10000]
	// 用于分页查询角色列表，最小值为1，最大值为10000
	Page int `json:"page" validate:"omitempty,gte=1,lte=10000" example:"1"`

	// PageSize 每页显示的角色数，选填，范围限制：[1,10000]
	// 用于限制每页显示角色的数量，最小值为1，最大值为10000
	PageSize int `json:"pageSize" validate:"omitempty,gte=1,lte=10000" example:"10"`

	// ID 角色ID，选填，UUID格式
	// 用于过滤查询特定角色ID的角色，格式必须为UUID4
	ID string `json:"id" validate:"omitempty,uuid4" example:"clywh0xv70001rvpgzd6256ns"`

	// Name 角色名称，选填，长度限制：3-128字符
	// 用于过滤查询特定角色名称的角色，最小长度为3，最大长度为128字符
	Name string `json:"name" validate:"omitempty,min=3,max=128" example:"admin"`

	// Status 角色状态，选填，1表示启用，0表示禁用
	// 用于过滤查询角色的启用/禁用状态，1表示启用，0表示禁用
	Status int `json:"status" validate:"omitempty,oneof=0 1" example:"1"`
}

// InfoRoleRequest 用于查询角色详情的请求体结构
type InfoRoleRequest struct {
	// ID 角色ID，必填，UUID格式
	// 唯一标识要删除的角色，格式必须为UUID4
	ID string `json:"id" validate:"required,uuid4" example:"clywh0xv70001rvpgzd6256ns"`
}

// AddRoleRequest 用于新增角色的请求体结构
type AddRoleRequest struct {
	// Name 角色名称，必填，长度限制：3-128字符
	// 角色名称必须唯一，建议使用字母数字组合，避免特殊字符，长度在3到128之间
	Name string `json:"name" validate:"required,min=3,max=128" example:"admin"`

	// Description 角色描述，选填，最大长度255字符
	// 用于对角色进行描述，长度限制在255字符内
	Description string `json:"description" validate:"omitempty,max=255" example:"Administrator role with full access"`

	// Status 角色状态，必填，1表示启用，0表示禁用
	// 如果未明确设置，则角色默认为启用
	Status int `json:"status" validate:"required,oneof=0 1" example:"1"`

	// Remark 备注，选填，最大长度256字符
	// 用于对角色进行描述或标记
	Remark string `json:"remark" validate:"omitempty,max=256" example:"This is a remark"`

	// PathIDList 路径ID列表，选填，用于指定该角色能够访问的路径
	// 角色可以访问多个路径，路径ID是与路径表中的路径关联的
	// 使用 "dive" 校验每个元素是否符合 UUID 格式
	PathIDList []string `json:"pathIDList" validate:"omitempty,dive,uuid4" example:"path_id_1,path_id_2"`
}

// EditRoleRequest 用于编辑角色信息的请求体结构
type EditRoleRequest struct {
	// ID 角色ID，必填，UUID格式
	// 用于唯一标识要编辑的角色，格式必须为UUID4
	ID string `json:"id" validate:"required,uuid4" example:"clywh0xv70001rvpgzd6256ns"`

	// Name 角色名称，选填，长度限制：3-128字符
	// 如果需要修改角色名称，则需要提供，且长度应为3到128字符
	Name string `json:"name" validate:"omitempty,min=3,max=128" example:"admin"`

	// Description 角色描述，选填，最大长度255字符
	// 如果需要修改角色描述，则需要提供，最大长度为255字符
	Description string `json:"description" validate:"omitempty,max=255" example:"Updated description"`

	// Status 角色状态，选填，1表示启用，0表示禁用
	// 如果需要修改角色的启用/禁用状态，则需要提供，1表示启用，0表示禁用
	Status int `json:"status" validate:"omitempty,oneof=0 1" example:"1"`

	// Remark 备注，选填，最大长度256字符
	// 用于对角色进行附加描述或标记
	Remark string `json:"remark" validate:"omitempty,max=256" example:"This is a remark"`

	// PathIDList 路径ID列表，选填，用于指定该角色能够访问的路径
	// 角色可以访问多个路径，路径ID是与路径表中的路径关联的
	// 使用 "dive" 校验每个元素是否符合 UUID 格式
	PathIDList []string `json:"pathIDList" validate:"omitempty,dive,uuid4" example:"path_id_1,path_id_2"`
}

// DelRoleRequest 用于删除角色的请求体结构
type DelRoleRequest struct {
	// ID 角色ID，必填，UUID格式
	// 唯一标识要删除的角色，格式必须为UUID4
	ID string `json:"id" validate:"required,uuid4" example:"clywh0xv70001rvpgzd6256ns"`
}

type ListRoleResponse struct {
	// Total 总条数
	Total int64 `json:"total" example:"100"`
	// List 数据
	List []RoleInfo `json:"list"`
}

// RoleInfo 包含角色的详细信息
type RoleInfo struct {
	// ID 角色ID，UUID格式
	// 用于唯一标识角色
	ID string `json:"id" example:"clywh0xv70001rvpgzd6256ns"`

	// Name 角色名称
	// 唯一标识角色的名称，3-128字符
	Name string `json:"name" example:"admin"`

	// Description 角色描述
	// 对角色的简短描述，最多255字符
	Description string `json:"description" example:"Administrator role with full access"`

	// Status 角色状态
	// 1表示启用，0表示禁用
	Status int `json:"status" example:"1"`

	// CreatedAt 角色创建时间
	// 格式为时间戳，标识角色的创建时间
	CreatedAt string `json:"createdAt" example:"2024-11-18T10:00:00Z"`

	// UpdatedAt 角色更新时间
	// 格式为时间戳，标识角色的最后更新时间
	UpdatedAt string `json:"updatedAt" example:"2024-11-18T11:00:00Z"`
}

type InfoRoleResponse struct {
	// ID 角色ID，UUID格式
	// 用于唯一标识角色
	ID string `json:"id" example:"clywh0xv70001rvpgzd6256ns"`

	// Name 角色名称
	// 唯一标识角色的名称，3-128字符
	Name string `json:"name" example:"admin"`

	// Description 角色描述
	// 对角色的简短描述，最多255字符
	Description string `json:"description" example:"Administrator role with full access"`

	// Status 角色状态
	// 1表示启用，0表示禁用
	Status int `json:"status" example:"1"`

	// CreatedAt 角色创建时间
	// 格式为时间戳，标识角色的创建时间
	CreatedAt string `json:"createdAt" example:"2024-11-18T10:00:00Z"`

	// UpdatedAt 角色更新时间
	// 格式为时间戳，标识角色的最后更新时间
	UpdatedAt string `json:"updatedAt" example:"2024-11-18T11:00:00Z"`

	// MenuData 是菜单树数据
	MenuData []*MenuNode `json:"menuData"`
}
