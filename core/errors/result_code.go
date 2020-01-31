package errors

var (
	// OK 成功
	OK = NewResultCode(0, "ok")

	// RequestError 请求错误
	RequestError = NewResultCode(400, "请求错误")

	// Unauthorized 认证错误
	Unauthorized = NewResultCode(401, "认证错误")

	// ForbiddenAccess 禁止访问
	ForbiddenAccess = NewResultCode(403, "禁止访问")

	// PathNotFound 请求地址不存在
	PathNotFound = NewResultCode(404, "请求地址不存在")

	// MethodNotAllowed 不支持该方法
	MethodNotAllowed = NewResultCode(405, "不支持该方法")

	// LoginExpires 登录失效
	LoginExpires = NewResultCode(450, "登录失效")

	// ServerError 服务器错误
	ServerError = NewResultCode(500, "服务器错误")

	// ServiceUnavailable 过载保护,服务暂不可用
	ServiceUnavailable = NewResultCode(503, "过载保护,服务暂不可用")

	// Deadline 服务调用超时
	Deadline = NewResultCode(504, "服务调用超时")

	// LimitExceed 超出限制
	LimitExceed = NewResultCode(509, "超出限制")

	// ParamError 参数错误
	ParamError = NewResultCode(600, "%s 参数错误")

	// FileTooLarge 文件过大
	FileTooLarge = NewResultCode(610, "文件过大")

	// FileTypeError 文件类型错误
	FileTypeError = NewResultCode(611, "文件类型错误")

	// FileNotExist 文件或目录不存在
	FileNotExist = NewResultCode(612, "文件或目录不存在")

	// FilePathIsNull 文件路径为空
	FilePathIsNull = NewResultCode(613, "文件路径为空")

	// FileReadFail 读取文件失败
	FileReadFail = NewResultCode(614, "读取文件失败")

	// ErrorUndefined 错误未定义
	ErrorUndefined = NewResultCode(996, "错误未定义")

	// BusinessFail 业务失败
	BusinessFail = NewResultCode(997, "业务失败")

	// SystemErr 系统异常
	SystemErr = NewResultCode(998, "系统异常")

	// UnknownError 未知错误
	UnknownError = NewResultCode(999, "未知错误")

	// LoadConfigFileFail 加载配置文件失败
	LoadConfigFileFail = NewResultCode(100000, "加载配置文件失败")

	// ObjectCopyFail 对象转换失败
	ObjectCopyFail = NewResultCode(100011, "对象转换失败")

	// ParseTimeFail 转换时间失败
	ParseTimeFail = NewResultCode(100012, "转换时间失败")

	// TemplateRenderFail 模板解析失败
	TemplateRenderFail = NewResultCode(100013, "模板解析失败")

	// EncryptDecryptFail 加解密失败
	EncryptDecryptFail = NewResultCode(100020, "加解密失败")

	// ObjectNotArray 对象不是数组
	ObjectNotArray = NewResultCode(100030, "对象不是数组")

	// JsonEncodeFail JSON序列化失败
	JsonEncodeFail = NewResultCode(100040, "JSON序列化失败")

	// JsonDecodeFail JSON反序列化失败
	JsonDecodeFail = NewResultCode(100041, "JSON反序列化失败")

	// DbOperationFail 数据库操作失败
	DbOperationFail = NewResultCode(100050, "数据库操作失败")

	// HttpRequestFail Http请求失败
	HttpRequestFail = NewResultCode(100060, "Http请求失败")

	// HttpCreateRequestFail Http创建请求失败
	HttpCreateRequestFail = NewResultCode(100061, "Http创建请求失败")

	// RedisOperationFail redis操作失败
	RedisOperationFail = NewResultCode(100101, "redis操作失败")

	// RedisConfigNotExist redis配置不存在
	RedisConfigNotExist = NewResultCode(100102, "redis配置不存在")

	// RedisNotInit redis配置不存在
	RedisNotInit = NewResultCode(100103, "redis未初始化")

	// RedisConnGetFail 获取redis连接失败
	RedisConnGetFail = NewResultCode(100104, "获取redis连接失败")

	// WX_OPERATION_ERROR 微信操作错误
	WX_OPERATION_ERROR = NewResultCode(100201, "%s 微信操作错误")
)
