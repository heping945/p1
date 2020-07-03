package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"pb/global"
	"pb/global/response"
	"pb/libs"

	"pb/model"
)

func CreateUser(c *gin.Context) {
	var user model.AuthenCode
	if err := c.ShouldBindJSON(&user); err != nil {
		response.FailWithMessage("参数错误", c)
	} else {
		user.Code = libs.HashAndSalt([]byte(user.Code))
		log.Println(user.Code)
		if err := global.GVA_DB.Create(&user).Error; err != nil {
			response.FailWithMessage("创建错误", c)
		} else {
			response.OkWithData(user, c)
		}
	}
}

func VerifyUser(c *gin.Context) {
	code := c.Param("code")
	if libs.VerifyCode(code) {
		response.OkWithMessage("验证正确", c)
	} else {
		response.FailWithMessage("验证错误", c)
	}
}
