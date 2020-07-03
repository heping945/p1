package middleware

import (
	"github.com/gin-gonic/gin"
	"pb/libs"
)

// 处理跨域请求,支持options访问

func Authen() gin.HandlerFunc {
	return func(c *gin.Context) {

		if h := c.Request.Header.Get("Authorization"); h == "" {
			c.JSON(401, gin.H{
				"code": 2004,
				"msg":  "未获得授权",
			})
			c.Abort()
		} else {
			if libs.VerifyCode(h) {
				c.Next()
			} else {
				c.JSON(401, gin.H{
					"code": 2004,
					"msg":  "授权码错误",
				})
				c.Abort()
			}
		}

	}
}
