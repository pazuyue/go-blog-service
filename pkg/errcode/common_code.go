package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的 AppKey 和 AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token 错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token 超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token 生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
	UnauthorizedAuthFail      = NewError(100009, "权限验证不通过")

	// 权限管理错误码
	ErrorCasbinCreateFail = NewError(601001, "创建权限规则失败")
	ErrorCasbinUpdateFail = NewError(601002, "更新权限规则失败")
	ErrorCasbinListFail   = NewError(601003, "获取权限规则列表失败")
)
