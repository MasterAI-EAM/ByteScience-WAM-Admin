package auth

// MenuTreeResponse 是返回菜单树的响应结构
type MenuTreeResponse struct {
	// Data 是菜单树的根节点数组
	Data []*MenuNode `json:"data"`
}

// MenuNode 是菜单节点，用于表示一个菜单及其子菜单和相关路径
type MenuNode struct {
	// ID 菜单ID
	ID string `json:"id" example:"clywh0xv70001rvpgzd6256ns"`

	// ParentID 父菜单ID，根菜单的 ParentID 为 null
	ParentID string `json:"parent_id" example:"null"`

	// Name 菜单名称
	Name string `json:"name" example:"Dashboard"`

	// CreatedAt 创建时间
	CreatedAt string `json:"createdAt" example:"2024-11-18T10:00:00Z"`

	// UpdatedAt 更新时间
	UpdatedAt string `json:"updatedAt" example:"2024-11-18T10:00:00Z"`

	// DeletedAt 软删除时间
	DeletedAt string `json:"deletedAt" example:"null"`

	// SubMenus 子菜单列表，包含所有直接的子菜单
	SubMenus []*MenuNode `json:"subMenus,omitempty"`

	// Paths 菜单下的路径信息
	Paths []*PathInfo `json:"paths,omitempty"`
}

// PathInfo 是菜单下的路径信息，用于表示与该菜单关联的接口
type PathInfo struct {
	// Path 路由路径
	Path string `json:"path" example:"/dashboard"`

	// Method HTTP方法（GET, POST, PUT, DELETE）
	Method string `json:"method" example:"GET"`

	// Description 路径描述
	Description string `json:"description" example:"Dashboard home page"`
}
