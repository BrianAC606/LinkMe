package req

type SignUpReq struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SMSReq struct {
	Number string `json:"number"`
}

type ChangeReq struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UsernameReq struct {
	Username string `json:"username"`
}

type DeleteUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UpdateProfileReq struct {
	RealName string `json:"realName"` // 真实姓名
	Avatar   string `json:"avatar"`   // 头像URL
	About    string `json:"about"`    // 个人简介
	Birthday string `json:"birthday"` // 生日
	Phone    string `json:"phone"`
}

type UpdateProfileAdminReq struct {
	UserID   int64  `json:"userId"`   // 用户ID
	RealName string `json:"realName"` // 真实姓名
	Avatar   string `json:"avatar"`   // 头像URL
	About    string `json:"about"`    // 个人简介
	Birthday string `json:"birthday"` // 生日
	Phone    string `json:"phone"`    // 手机号
}

type LoginSMSReq struct {
	Code string `json:"code"`
}

type ListUserReq struct {
	Page int    `json:"page,omitempty"` // 当前页码
	Size *int64 `json:"size,omitempty"` // 每页数据量
}

type GetUserCountReq struct {
}

type LogoutReq struct {
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refreshToken"`
}

type GetProfileReq struct {
}
