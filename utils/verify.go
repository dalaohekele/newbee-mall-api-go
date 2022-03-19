package utils

var (
	GoodsCategoryVerify           = Rules{"CategoryRank": {NotEmpty()}, "CategoryName": {NotEmpty()}}
	AdminUserRegisterVerify       = Rules{"Username": {NotEmpty()}, "NickName": {NotEmpty()}, "Password": {NotEmpty()}}
	AdminUserChangePasswordVerify = Rules{"Password": {NotEmpty()}}
)
