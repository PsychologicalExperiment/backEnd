package errorcode

import (
	"fmt"

	errCodePb "github.com/PsychologicalExperiment/backEnd/api/error_code"
)

const (
	OKCode uint32 = iota + 0
	ErrParamsTypeErrorInServer
	ErrParamsInvalid
	ErrQueryRecordNotFound
	ErrRearcherNotExist
	ErrInvalidResearcher
)

// 对外错误码
var errMsgMap = map[uint32]string{
	OKCode:                     "ok",
	ErrParamsTypeErrorInServer: "internal error, param wrong in server",
	ErrParamsInvalid:           "invalid input params",
	ErrQueryRecordNotFound:     "the experiment not found",
	ErrRearcherNotExist:        "the user does not exist",
	ErrInvalidResearcher:       "the user_id is invalid",
}

// 对外错误信息
var errToCommonCode = map[uint32]uint32{
	OKCode:                     uint32(errCodePb.ErrorCode_CODE_SUCC),
	ErrParamsTypeErrorInServer: uint32(errCodePb.ErrorCode_CODE_INTERNAL_ERR),
	ErrParamsInvalid:           uint32(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrQueryRecordNotFound:     uint32(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrRearcherNotExist:        uint32(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrInvalidResearcher:       uint32(errCodePb.ErrorCode_CODE_PARAM_ERR),
}

type ErrorImpl struct {
	ErrorCode uint32
	ErrorMsg  string
}

func New(code uint32) ErrorImpl {
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
