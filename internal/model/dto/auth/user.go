package auth

type AddRequest struct {
}

type EditRequest struct {
}

type DelRequest struct {
}

type ListRequest struct {
	// page 页码 参数必填 int  参数范围：[1,-]
	Page int `form:"page" validate:"required,gte=1,lte=10000" example:"1"`
	// pageSize 页数 参数必填 int  参数范围：[1,-]
	PageSize int `form:"pageSize" validate:"required,gte=1,lte=10000" example:"10"`
}

type ListResponse struct {
	// total 总条数
	Total int64 `json:"total" example:"100"`
	// data 数据
	Data []*UserInfo `json:"data"`
}

type UserInfo struct {
	// id string 编号
	ID string `json:"id" example:"clywh0xv70001rvpgzd6256ns"`
	// name string  用户名
	Name string `json:"name" example:"user_name"`
}
