package request

type MallAdminLoginParam struct {
	UserName    string `json:"userName"`
	PasswordMd5 string `json:"passwordMd5"`
}
type MallAdminParam struct {
	LoginUserName string `json:"loginUserName"`
	LoginPassword string `json:"loginPassword"`
	NickName      string `json:"nickName"`
}
