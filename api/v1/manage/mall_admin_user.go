package manage

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common/request"
	"main.go/model/common/response"
	"main.go/model/manage"
	manageReq "main.go/model/manage/request"
	"main.go/utils"
	"strconv"
)

type MallAdminUserApi struct {
}

// CreateMallAdminUser 创建MallAdminUser
func (mallAdminUserApi *MallAdminUserApi) CreateMallAdminUser(c *gin.Context) {
	var params manageReq.MallAdminParam
	_ = c.ShouldBindJSON(&params)
	if err := utils.Verify(params, utils.AdminUserRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mallAdminUser := manage.MallAdminUser{
		LoginUserName: params.LoginUserName,
		NickName:      params.NickName,
		LoginPassword: utils.MD5V([]byte(params.LoginPassword)),
	}
	if err := mallAdminUserService.CreateMallAdminUser(mallAdminUser); err != nil {
		global.GVA_LOG.Error("创建失败:", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// UpdateMallAdminUserPassword 修改密码
func (mallAdminUserApi *MallAdminUserApi) UpdateMallAdminUserPassword(c *gin.Context) {
	var params manage.MallAdminUser
	_ = c.ShouldBindJSON(&params)
	if err := utils.Verify(params, utils.AdminUserChangePasswordVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mallAdminUserName := manage.MallAdminUser{
		LoginPassword: utils.MD5V([]byte(params.LoginPassword)),
	}
	if err := mallAdminUserService.UpdateMallAdminUser(mallAdminUserName); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}

}

// UpdateMallAdminUserName 更新MallAdminUser用户名
func (mallAdminUserApi *MallAdminUserApi) UpdateMallAdminUserName(c *gin.Context) {
	var mallAdminUser manage.MallAdminUser
	_ = c.ShouldBindJSON(&mallAdminUser)
	if err := mallAdminUserService.UpdateMallAdminUser(mallAdminUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// AdminUserProfile 用id查询AdminUser
func (mallAdminUserApi *MallAdminUserApi) AdminUserProfile(c *gin.Context) {
	adminToken := c.GetHeader("token")
	if err, mallAdminUser := mallAdminUserService.GetMallAdminUser(adminToken); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("未查询到记录", c)
	} else {
		mallAdminUser.LoginPassword = "******"
		response.OkWithData(mallAdminUser, c)
	}
}

// AdminLogin 管理员登陆
func (mallAdminUserApi *MallAdminUserApi) AdminLogin(c *gin.Context) {
	var adminLoginParams manageReq.MallAdminLoginParam
	_ = c.ShouldBindJSON(&adminLoginParams)
	if err, _, adminToken := mallAdminUserService.AdminLogin(adminLoginParams); err != nil {
		response.FailWithMessage("登陆失败", c)
	} else {
		response.OkWithData(adminToken.Token, c)
	}
}

// AdminLogout 登出
func (mallAdminUserApi *MallAdminUserApi) AdminLogout(c *gin.Context) {
	var adminUser manage.MallAdminUser
	_ = c.ShouldBindJSON(adminUser)
	var ids request.IdsReq
	ids.Ids = append(ids.Ids, int(adminUser.AdminUserId))
	if err := mallAdminUserTokenService.DeleteMallAdminUserTokenByIds(ids); err != nil {
		response.FailWithMessage("登出失败", c)
	} else {
		response.OkWithMessage("登出成功", c)
	}

}

// UserList 商城注册用户列表
func (mallAdminUserApi *MallAdminUserApi) UserList(c *gin.Context) {
	var pageInfo manageReq.MallUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallUserService.GetMallUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// LockUser 用户禁用与解除禁用(0-未锁定 1-已锁定)
func (mallAdminUserApi *MallAdminUserApi) LockUser(c *gin.Context) {
	lockStatus, _ := strconv.Atoi(c.Param("lockStatus"))
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := mallUserService.LockUser(IDS, lockStatus); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
