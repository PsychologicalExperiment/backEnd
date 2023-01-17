package serverErr

import (
	"fmt"

	errCodePb "github.com/PsychologicalExperiment/backEnd/api/error_code"
	userInfoPb "github.com/PsychologicalExperiment/backEnd/api/user_info_server"
	"google.golang.org/grpc/grpclog"
)

const (
	OKCode ErrCode = iota + 0
	// 注册相关
	ErrUserInfoNotProvided
	ErrParamsTypeErrorInServer
	ErrMySqlError
	ErrEmailAlreadyUsed
	ErrEmailNotProvided
	ErrPasswordNotProvided
	ErrGenderInvalid
	ErrUserTypeInvalid
	ErrPhoneNumberAlreadyUsed
	ErrGenerateTokenFailed
	ErrSetRedisFailed
	// 登录相关
	ErrorUserNotFound
	ErrorPasswordNotRight
)

var errMsgMap = map[ErrCode]string{
	OKCode: "ok",
	// 注册相关
	ErrUserInfoNotProvided:     "param error, user info is not provided",
	ErrParamsTypeErrorInServer: "internal error, param wrong in server",
	ErrMySqlError:              "internal error, db sql error, please contact us",
	ErrEmailAlreadyUsed:        "param error, email is already used",
	ErrEmailNotProvided:        "param error, email is empty",
	ErrPasswordNotProvided:     "param error, password is empty",
	ErrGenderInvalid:           "param error, gender is invalid",
	ErrUserTypeInvalid:         "param error, user type is invalid",
	ErrPhoneNumberAlreadyUsed:  "param error, phone number is already used",
	ErrGenerateTokenFailed:     "internal error, generate token failed",
	ErrSetRedisFailed:          "internal error, set redis failed",
	// 登录相关
	ErrorUserNotFound:     "param error, user not found",
	ErrorPasswordNotRight: "param error, password not correct",
}

var errToCommonCode = map[ErrCode]ErrCode{
	OKCode: ErrCode(errCodePb.ErrorCode_CODE_SUCC),
	// 注册相关
	ErrUserInfoNotProvided:     ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrParamsTypeErrorInServer: ErrCode(errCodePb.ErrorCode_CODE_INTERNAL_ERR),
	ErrMySqlError:              ErrCode(errCodePb.ErrorCode_CODE_INTERNAL_ERR),
	ErrEmailAlreadyUsed:        ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrEmailNotProvided:        ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrPasswordNotProvided:     ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrGenderInvalid:           ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrUserTypeInvalid:         ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrPhoneNumberAlreadyUsed:  ErrCode(errCodePb.ErrorCode_CODE_PARAM_ERR),
	ErrGenerateTokenFailed:     ErrCode(errCodePb.ErrorCode_CODE_INTERNAL_ERR),
	ErrSetRedisFailed:          ErrCode(errCodePb.ErrorCode_CODE_INTERNAL_ERR),
	// 登录相关
	ErrorUserNotFound:     ErrCode(errCodePb.ErrorCode_CODE_INTERNAL_ERR),
	ErrorPasswordNotRight: ErrCode(errCodePb.ErrorCode_CODE_INTERNAL_ERR),
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
		CommonRsp: &userInfoPb.CommonRsp{
			Code: uint32(myerr.ErrorCode),
			Msg:  myerr.ErrorMsg,
		},
	}
}

func CommonRsp(
	err error,
) *userInfoPb.CommonRsp {
	myerr, ok := err.(ErrorImpl)
	if !ok {
		grpclog.Errorf("ErrRspQueryFromDruid|error usage of error code")
		myerr = New(ErrParamsTypeErrorInServer)
	}
	return &userInfoPb.CommonRsp{
		Code: uint32(myerr.ErrorCode),
		Msg:  myerr.ErrorMsg,
	}
}
