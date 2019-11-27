package errors

var (
	//成功
	OK = AddResultCode(0, "OK", "ResultCode.OK")

	//token错误
	RequestError = AddResultCode(400, "请求错误", "ResultCode.RequestError")

	//认证错误
	Unauthorized = AddResultCode(401, "认证错误", "ResultCode.Unauthorized")

	//禁止访问
	ForbiddenAccess = AddResultCode(403, "禁止访问", "ResultCode.ForbiddenAccess")

	//请求地址不存在
	PathNotFound = AddResultCode(404, "请求地址不存在", "ResultCode.PathNotFound")

	//不支持该方法
	MethodNotAllowed = AddResultCode(405, "不支持该方法", "ResultCode.MethodNotAllowed")

	//Token过期
	TokenExpires = AddResultCode(450, "登录失效", "ResultCode.TokenExpires")

	//请求参数错误
	ServerError = AddResultCode(500, "服务器错误", "ResultCode.ServerError")

	//过载保护,服务暂不可用
	ServiceUnavailable = AddResultCode(503, "过载保护,服务暂不可用", "ResultCode.ServiceUnavailable")

	//服务调用超时
	Deadline = AddResultCode(504, "服务调用超时", "ResultCode.Deadline")

	//超出限制
	LimitExceed = AddResultCode(509, "超出限制", "ResultCode.LimitExceed")

	//参数错误
	ParamError = AddResultCode(600, "参数错误", "ResultCode.ParamError")

	//文件过大
	FileTooLarge = AddResultCode(610, "文件过大", "ResultCode.FileTooLarge")

	//文件类型错误
	FileTypeError = AddResultCode(611, "文件类型错误", "ResultCode.FileTypeError")

	//文件或目录不存在
	FileNotExist = AddResultCode(612, "文件或目录不存在", "ResultCode.FileNotExist")

	//文件路径为空
	FilePathIsNull = AddResultCode(613, "文件路径为空", "ResultCode.FilePathIsNull")

	//读取文件失败
	FileReadFail = AddResultCode(614, "读取文件失败", "ResultCode.FileReadFail")

	//错误未定义
	ErrorUndefined = AddResultCode(996, "错误未定义", "ResultCode.ErrorUndefined")

	//业务失败
	BusinessFail = AddResultCode(997, "业务失败", "ResultCode.BusinessFail")

	//系统异常
	SystemErr = AddResultCode(998, "系统异常", "ResultCode.SystemErr")

	//未知错误
	UnknownError = AddResultCode(999, "未知错误", "ResultCode.UnknownError")

	RocketMQProduceInitError   = AddResultCode(10001, "RocketMQ Produce 初始化异常", "ResultCode.RocketMQProduceInitError")
	RocketMQSendMsgError       = AddResultCode(10002, "RocketMQ SendMsg 失败", "ResultCode.RocketMQSendMsgError")
	RocketMQConsumerInitError  = AddResultCode(10003, "RocketMQ Consumer 初始化异常", "ResultCode.RocketMQConsumerInitError")
	RocketMQConsumerStartError = AddResultCode(10004, "RocketMQ Consumer 启动异常", "ResultCode.RocketMQConsumerStartError")
	RocketMQConsumerStopError  = AddResultCode(10005, "RocketMQ Consumer 停止异常", "ResultCode.RocketMQConsumerStopError")

	KafkaMqSendMsgError           = AddResultCode(10101, "Kafka发送消息失败", "ResultCode.KafkaMqSendMsgError")
	KafkaMqSendMsgCantBeNullError = AddResultCode(10102, "Kafka发送的消息不能为空", "ResultCode.KafkaMqSendMsgCantBeNullError")
	KafkaMqConsumeMsgError        = AddResultCode(10103, "Kafka消费消息失败", "ResultCode.KafkaMqConsumeMsgError")
	KafkaMqConsumeStartError      = AddResultCode(10104, "Kafka消费启动失败", "ResultCode.KafkaMqConsumeStartError")

	TryDistributedLockError = AddResultCode(10201, "获取分布式锁异常", "ResultCode.TryDistributedLockError")
	GetDistributedLockError = AddResultCode(10202, "获取分布式锁失败", "ResultCode.GetDistributedLockError")
	MysqlOperateError       = AddResultCode(10203, "db操作出现异常", "ResultCode.MysqlOperateError")
	RedisOperateError       = AddResultCode(10204, "redis操作出现异常", "ResultCode.RedisOperateError")

	DbMQSendMsgError         = AddResultCode(10301, "Db 保存 message queue 失败", "ResultCode.DbMQSendMsgError")
	DbMQCreateConsumerError  = AddResultCode(10302, "Db 创建 message queue consumer 失败", "ResultCode.DbMQCreateConsumerError")
	DbMQConsumerStartedError = AddResultCode(10303, "Db 创建 message queue consumer 已启动", "ResultCode.DbMQConsumerStartedError")
)
