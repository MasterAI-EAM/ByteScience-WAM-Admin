package auth

// MenuTreeResponse 是返回菜单树的响应结构
type MenuTreeResponse struct {
	// Data 是菜单树的根节点数组
	Data []*MenuNode `json:"data"`
}

// BaseNode 定义菜单和路径的公共字段
type BaseNode struct {
	// ID 节点ID
	ID string `json:"id" example:"clywh0xv70001rvpgzd6256ns"`

	// ParentID 父节点ID，根节点的 ParentID 为 null
	ParentID string `json:"parent_id" example:"null"`

	// Name 节点名称
	Name string `json:"name" example:"Dashboard"`
}

// MenuNode 表示菜单及其子菜单
type MenuNode struct {
	BaseNode

	// MenuData 子菜单列表
	MenuData []*MenuNode `json:"menuData,omitempty"`

	// Paths 菜单下的路径信息
	Paths []*PathInfo `json:"paths,omitempty"`
}

// PathInfo 表示菜单下的路径信息
type PathInfo struct {
	// Path 路由路径
	Path string `json:"path" example:"/dashboard"`

	// Method HTTP方法（GET, POST, PUT, DELETE）
	Method string `json:"method" example:"GET"`

	// Description 路径描述
	Description string `json:"description" example:"Dashboard home page"`
}
