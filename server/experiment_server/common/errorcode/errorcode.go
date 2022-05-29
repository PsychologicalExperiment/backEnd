package errorcode

import (
	"fmt"

	errCodePb "github.com/PsychologicalExperiment/backEnd/api/error_code"
)

const (
	OKCode ErrCode = iota + 0
	ErrParamsTypeErrorInServer
	ErrParamsInvalid
	ErrQueryRecordNotFound
)

//  对外错误码
var errMsgMap = map[ErrCode]string{
	OKCode:                     "ok",
	ErrParamsTypeErrorInServer: "internal error, param wrong in server",
	ErrParamsInvalid:           "invalid input params",
	ErrQueryRecordNotFound:     "the experiment not found",
}

//  对外错误信息
var errToCommonCode = map[ErrCode]ErrCode{
	OKCode:                     ErrCode(errCodePb.ErrorCode_CODE_SUCC),
	ErrParamsTypeErrorInServer: ErrCode(errCodePb.ErrorCode_CODE_INTERNAL_ERR),
	ErrParamsInvalid:           ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrQueryRecordNotFound:     ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
}

type ErrCode int32

type ErrorImpl struct {
	ErrorCode ErrCode
	ErrorMsg  string
}

func New(code ErrCode) ErrorImpl {
	return ErrorImpl{
		ErrorCode: errToCommonCode[code],
		ErrorMsg:  errMsgMap[code],
	}
}

func (e ErrorImpl) Error() string {
	strFormat := `
	errorCode: %d
	errorMsg: %s
	`
	return fmt.Sprintf(strFormat, e.ErrorCode, e.ErrorMsg)
}
