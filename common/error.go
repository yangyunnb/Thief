package common

type ErrorCode int
type Error struct {
	ErrCode     ErrorCode
	ErrMsg      string
	InnerErrMag string
}

const (
	ESuccess      = 0
	EParamInvalid = 10001
	EDefault      = 99999
)

var errorMap = map[ErrorCode]Error{
	ESuccess:      {ESuccess, "success", ""},
	EParamInvalid: {ESuccess, "param invalid", ""},
	EDefault:      {EDefault, "default error", "default error"},
}

func GetErrorByErrCode(code ErrorCode, innerMsg string) Error {
	if err, ok := errorMap[code]; ok {
		if len(innerMsg) > 0 {
			err.InnerErrMag = innerMsg
		}
		return err
	}
	return errorMap[EDefault]
}
