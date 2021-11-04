package terror

type MsgBase struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Msg struct {
	MsgBase
	Data interface{} `json:"data"`
}

var (
	Succeed = MsgBase{
		Code: 0,
		Msg:  "succeed",
	}

	ErrInvalidParam = MsgBase{
		Code: 101,
		Msg:  "invalid param",
	}

	ErrUnexpected = MsgBase{
		Code: 102,
		Msg:  "unexpected internal error",
	}

	// 200
	ErrDataNotFound = MsgBase{
		Code: 201,
		Msg:  "data not found",
	}

	ErrDuplicateUser = MsgBase{
		Code: 202,
		Msg:  "user duplicate error",
	}

	ErrInvalidToken = MsgBase{
		Code: 203,
		Msg:  "token invalid",
	}
)
