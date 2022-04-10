package serverErr

import (
	"fmt"

	commonPb "github.com/PsychologicalExperiment/backEnd/api/api_common"
	errCodePb "github.com/PsychologicalExperiment/backEnd/api/error_code"
	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	"google.golang.org/grpc/grpclog"
)

const (
	OKCode ErrCode = iota + 0
	ErrUserInfoNotProvided
	ErrParamsTypeErrorInServer
	ErrMySqlError
	ErrEmailAlreadyUsed
)

var errMsgMap = map[ErrCode]string{
	OKCode:                     "ok",
	ErrUserInfoNotProvided:     "param error, user info is not provided",
	ErrParamsTypeErrorInServer: "internal error, param wrong in server",
	ErrMySqlError:              "internal error, db sql error, please contact us",
	ErrEmailAlreadyUsed:        "param error, email already used",
}

var errToCommonCode = map[ErrCode]ErrCode{
	OKCode:                     ErrCode(errCodePb.ErrorCode_CODE_SUCC),
	ErrUserInfoNotProvided:     ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrParamsTypeErrorInServer: ErrCode(errCodePb.ErrorCode_CODE_INTERNAL_ERR),
	ErrMySqlError:              ErrCode(errCodePb.ErrorCode_CODE_INTERNAL_ERR),
	ErrEmailAlreadyUsed:        ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
}

type ErrCode uint32

// ErrorImpl
type ErrorImpl struct {
	ErrorCode ErrCode
	ErrorMsg  string
}

// New
func New(code ErrCode) ErrorImpl {
	return ErrorImpl{
		ErrorCode: errToCommonCode[code],
		ErrorMsg:  errMsgMap[code],
	}
}

// Error
func (e ErrorImpl) Error() string {
	strFormat := `
    Error in api_data_query_server
    errorCode: %d
    errorMsg: %s
	`
	return fmt.Sprintf(strFormat, e.ErrorCode, e.ErrorMsg)
}

func RegisterErrRsp(
	err error,
) *userInfoPb.RegisterRsp {
	myerr, ok := err.(ErrorImpl)
	if !ok {
		grpclog.Errorf("ErrRspQueryFromDruid|error usage of error code")
		myerr = New(ErrParamsTypeErrorInServer)
	}
	return &userInfoPb.RegisterRsp{
		CommonRsp: &commonPb.CommonRsp{
			Code: uint32(myerr.ErrorCode),
			Msg:  myerr.ErrorMsg,
		},
		Token: "",
	}
}
