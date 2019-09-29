package common

//标准错误码
const (
	// 0:成功
	SUCC = iota
	// 1:失败
	ERR
	// 2:认证失败
	AUTH
	// 3:请求异常
	REQUEST
	// 4:参数异常
	PARAM
	// 5:多个结果
	MULTIPLE
	// 6:空结果
	EMPTY
)

const (
	SUCC_STR     = "请求成功"
	AUTH_STR     = "认证错误"
	REQUEST_STR  = "请求方式错误"
	PARAM_STR    = "参数格式化错误"
	SQL_STR      = "SQL错误"
	ROW_STR      = "转化行错误"
	PASSWORD_STR = "密码错误"
	LOCK_STR     = "用户已锁定"
	MULTIPLE_STR = "多个结果"
	EMPTY_STR    = "结果为空"
	JSON_STR     = "JSON转化错误"
)
