package auth

// AddAdminRequest 用于添加管理员的请求体结构
type AddAdminRequest struct {
	// UserName 用户名，必填，长度限制为3-128字符
	// 用户名必须唯一，且避免特殊字符，长度需满足最小值3字符，最大值128字符
	UserName string `json:"userName" validate:"required,min=3,max=128" example:"user1"`

	// Nickname 昵称，选填，最大长度128字符
	// 昵称是管理员的可选名称，长度限制为128字符
	Nickname string `json:"nickname" validate:"omitempty,max=128" example:"AdminNickname"`

	// Password 密码，必填，长度限制
	// 密码，且避免特殊字符，长度需满足最小值6字符，最大值32字符
	Password string `json:"password" validate:"required,min=6,max=32" example:"oldpassword123"`

	// Email 邮箱，选填，必须符合邮箱格式
	// 邮箱字段为可选项，但如果提供，必须符合标准的邮箱格式
	Email string `json:"email" validate:"omitempty,email" example:"user@example.com"`

	// Phone 手机号码，选填，必须符合E.164国际电话格式
	// 手机号码格式应为国际标准格式，例如 +1234567890
	Phone string `json:"phone" validate:"omitempty,e164" example:"+1234567890"`

	// Remark 备注，选填，最大长度256字符
	// 用于提供对管理员的附加说明或备注信息，最大长度为256字符
	Remark string `json:"remark" validate:"max=256" example:"This is a remark"`
}

// EditAdminRequest 用于编辑管理员信息的请求体结构
type EditAdminRequest struct {
	// ID 编号，必填，UUID格式
	// 用于唯一标识管理员的ID，格式必须符合UUID4标准
	ID string `json:"id" validate:"required,uuid4" example:"clywh0xv70001rvpgzd6256ns"`

	// UserName 用户名，选填，长度限制为3-128字符
	// 用户名必须唯一，且避免特殊字符，长度需满足最小值3字符，最大值128字符
	UserName string `json:"userName" validate:"omitempty,min=3,max=128" example:"user1"`

	// Nickname 昵称，选填，最大长度128字符
	// 昵称是管理员的可选名称，长度限制为128字符
	Nickname string `json:"nickname" validate:"omitempty,max=128" example:"AdminNickname"`

	// Email 邮箱，选填，必须符合邮箱格式
	// 仅当管理员更改邮箱时需要传递，且邮箱格式必须符合标准格式
	Email string `json:"email" validate:"omitempty,email" example:"user@example.com"`

	// Phone 手机号码，选填，必须符合E.164国际电话格式
	// 仅当管理员更改电话号码时需要传递，且格式应符合国际电话规范
	Phone string `json:"phone" validate:"omitempty,e164" example:"+1234567890"`

	// Remark 备注，选填，最大长度256字符
	// 备注用于对管理员的附加描述或标记，最大长度为256字符
	Remark string `json:"remark" validate:"max=256" example:"This is a remark"`
}

// DelAdminRequest 用于删除管理员的请求体结构
type DelAdminRequest struct {
	// ID 编号，必填，UUID格式
	// 唯一标识要删除的管理员，格式必须为UUID4
	ID string `json:"id" validate:"required,uuid4" example:"clywh0xv70001rvpgzd6256ns"`
}

// ListAdminRequest 用于查询管理员列表的请求体结构
type ListAdminRequest struct {
	// Page 页码，选填，范围限制：[1,10000]
	// 用于分页查询管理员列表，最小值为1，最大值为10000
	Page int `json:"page" validate:"omitempty,gte=1,lte=10000" example:"1"`

	// PageSize 每页大小，选填，范围限制：[1,10000]
	// 用于限制每页返回的管理员数量，最小值为1，最大值为10000
	PageSize int `json:"pageSize" validate:"omitempty,gte=1,lte=10000" example:"10"`

	// ID 编号，选填，UUID格式
	// 用于过滤查询特定ID的管理员，格式必须为UUID4
	ID string `json:"id" validate:"omitempty,uuid4" example:"clywh0xv70001rvpgzd6256ns"`

	// UserName 用户名，选填，长度限制：3-128字符
	// 用于过滤查询特定用户名的管理员，最小长度为3，最大长度为128字符
	UserName string `json:"userName" validate:"omitempty,min=3,max=128" example:"user1"`

	// Email 邮箱，选填，格式验证
	// 用于过滤查询特定邮箱的管理员，邮箱格式必须合法
	Email string `json:"email" validate:"omitempty,email" example:"user@example.com"`

	// Phone 手机号码，选填，E.164格式
	// 用于过滤查询特定手机号码的管理员，手机号必须符合国际标准E.164格式
	Phone string `json:"phone" validate:"omitempty,e164" example:"+1234567890"`
}

type ListAdminResponse struct {
	// total 总条数
	Total int64 `json:"total" example:"100"`
	// List 数据
	List []AdminInfo `json:"list"`
}

type AdminInfo struct {
	// ID string 编号
	ID string `json:"id" example:"clywh0xv70001rvpgzd6256ns"`
	// UserName string 用户名
	UserName string `json:"userName" example:"user1"`
	// Nickname string 昵称
	Nickname string `json:"nickname" example:"AdminNickname"`
	// Email string 邮箱
	Email string `json:"email" example:"user@example.com"`
	// Phone string 手机号码
	Phone string `json:"phone" example:"+1234567890"`
	// Remark string 备注
	Remark string `json:"remark" example:"This is a remark"`
	// LastLoginAt string 上次登录时间
	LastLoginAt string `json:"lastLoginAt" example:"2024-11-18T15:04:05Z"`
	// CreatedAt string 创建时间
	CreatedAt string `json:"createdAt" example:"2024-11-18T10:00:00Z"`
	// UpdatedAt string 更新时间
	UpdatedAt string `json:"updatedAt" example:"2024-11-18T11:00:00Z"`
}
