package dto

// UserRegisterRequest 用户注册请求
type UserRegisterRequest struct {
	NickName string `json:"nickName"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Key      string `json:"key"`
}

// UserLoginRequest 用户登录请求
type UserLoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

// SendEmailRequest 发送邮箱请求
type SendEmailRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	// 操作类型：1.绑定邮箱 2.解绑邮箱 3. 改密码
	OperationType uint `json:"operationType"`
}
