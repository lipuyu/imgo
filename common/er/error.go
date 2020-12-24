package er

var (
	ParamsError = Error{Msg:"传入参数不对", Code: 1}
	JsonBodyError = Error{Msg:"传入body格式不对", Code: 2}
	AppError = Error{Msg:"App名字不对", Code: 3}

	SourceNotExistError = Error{Msg:"资源不存在", Code: 404}

	UserNotExistError = Error{Msg:"用户不存在", Code: 1000}
	PasswordError = Error{Msg:"用户密码不对", Code: 1001}
	UserNameRepeatError = Error{Msg:"用户名已经存在", Code: 1002}
)

type Error struct {
	Msg string `json:"msg"`
	Code int `json:"code"`
}
