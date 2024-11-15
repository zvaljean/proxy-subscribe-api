package errno

var (
	OK                  = &BizCode{Code: 0, Msg: "ok"}
	InternalServerError = &BizCode{Code: 1, Msg: "internal server error"}
	ErrParam            = &BizCode{Code: 2, Msg: "request param error"}
	ErrTokenInvalid     = &BizCode{Code: 3, Msg: "the token was invalid"}
	ErrAccount          = &BizCode{Code: 4, Msg: "account or password error"}
	ErrPwdHasHan        = &BizCode{Code: 5, Msg: "pwd has han"}
	ErrNewPwdLen        = &BizCode{Code: 6, Msg: "the len of new pass should >= 8"}
	ErrDataExist        = &BizCode{Code: 7, Msg: "data existed"}
	ErrDataNotExist     = &BizCode{Code: 8, Msg: "data not exist"}
	ErrDBQuery          = &BizCode{Code: 9, Msg: "db query error"}
)

type BizCode struct {
	Code int
	Msg  string
	Err  error
}

func (err *BizCode) Error() string {
	return err.Msg
}

// type Err struct {
// 	BizCode
// }

// func (err *Err) Error() string {
// 	return fmt.Sprintf("Err - code: %d, msg: %s, error: %s", err.Code, err.Msg, err.Err)
// }

// func DecodeErr(err error) (int, string) {
// 	if err == nil {
// 		return OK.Code, OK.Msg
// 	}
//
// 	switch typed := err.(type) {
// 	case *Err:
// 		return typed.Code, typed.Msg
// 	case *BizCode:
// 		return typed.Code, typed.Msg
// 	default:
// 	}
// 	return InternalServerError.Code, err.Error()
// }
