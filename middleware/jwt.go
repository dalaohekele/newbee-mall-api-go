package middleware

import (
	"github.com/gin-gonic/gin"
	"main.go/model/common/response"
	"main.go/service"
	"time"
)

var manageAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserTokenService

func AdminJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		err, mallAdminUserToken := manageAdminUserTokenService.ExistAdminToken(token)
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if time.Now().After(mallAdminUserToken.ExpireTime) {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		c.Next()
	}

}
