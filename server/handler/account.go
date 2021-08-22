package handler

type requestUsernamePassword struct {
	Username string `json:"username" binding:"min=1,max=20" example:"username"` // 用户名
	Password string `json:"password" binding:"min=1,max=20" example:"password"` // 密码
}

func AddAccount() {

}

type requestAccountIds struct {
	Ids []int64 `json:"ids"`
}

func DeleteAccounts() {

}

func Login() {

}

type requestUpdateCurrentAccount struct {
	OldPassword string `json:"old_password" binding:"min=1,max=20" example:"password"` // 旧密码
	NewPassword string `json:"new_password" binding:"min=1,max=20" example:"password"` // 新密码
}

func UpdateCurrentAccountPassword() {

}

type requestUpdateAccount struct {
	Id          int64  `json:"id" binding:"gte=2" example:"2"`                         // 账户ID
	NewPassword string `json:"new_password" binding:"min=1,max=20" example:"password"` // 新密码
}

func UpdateAccountPassword() {

}

type responseAccount struct {
	Id        int64  `json:"id" example:"2"`                           // 账户ID
	Username  string `json:"username" example:"username"`              // 用户名
	LastLogin Time   `json:"last_login" example:"2006-01-02 15:04:05"` // 上次登陆时间
}

type responseFindAccounts struct {
	Code     int             `json:"code" example:"2"`   // 错误码，0为成功、其余为错误
	Count    int64           `json:"count" example:"10"` // 结果总数
	Accounts responseAccount `json:"accounts"`           // 结果
}

func FindAccounts() {

}

func GetCurrentAccount() {

}

func Logout() {

}
