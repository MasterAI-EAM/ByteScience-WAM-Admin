package utils

// 错误码定义
const (
	Success       = 0   // 成功
	BadRequest    = 400 // 请求错误
	InternalError = 500 // 服务器内部错误

	UserAlreadyExistsCode      = 1001 // 用户已存在
	UserNotFoundCode           = 1002 // 用户未找到
	UserInvalidCredentialsCode = 1003 // 用户凭证无效

)

// ErrorMessages 错误信息映射
var ErrorMessages = map[int]string{
	Success:       "success",
	BadRequest:    "Invalid Request Parameters",
	InternalError: "Internal Server Error",

	// 用户模块
	UserAlreadyExistsCode:      "User already exists",
	UserNotFoundCode:           "User not found",
	UserInvalidCredentialsCode: "Invalid credentials",
}
