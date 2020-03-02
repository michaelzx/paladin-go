package errs

var (
	Common           = NewBadRequestErrorWithCode(10000, "")
	DataNotExist     = NewBadRequestErrorWithCode(10001, "不存在")
	DataAlreadyExist = NewBadRequestErrorWithCode(10002, "已存在")
	ParamsNotExist   = NewBadRequestErrorWithCode(10003, "缺少参数：")
	ParamsErr        = NewBadRequestErrorWithCode(10004, "参数错误：")
	ParamsRequired   = NewBadRequestErrorWithCode(10005, "缺少必填项：")
)
