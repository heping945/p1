package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pb/global"
	"pb/global/response"
	"pb/model"
)

// 创建一个分类
func CreatePostCategory(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {

		response.FailWithMessage(fmt.Sprintf("发生了错误:%s", err), c)
	} else {
		if err := global.GVA_DB.Create(&category).Error; err != nil {
			response.FailWithMessage(fmt.Sprintf("failure 创建失败 name唯一"), c)
		} else {
			response.OkWithData(category, c)
		}
	}
}

func GetPostCategoryList(c *gin.Context) {
	var category []model.Category
	global.GVA_DB.Find(&category)

	response.OkWithData(category, c)
}

// 删除
func DeletePostCategory(c *gin.Context) {
	name := c.Param("name")
	global.GVA_DB.Where("name = ?", name).Unscoped().Delete(&model.Category{})
	response.OkWithMessage("删除成功", c)
}

// 修改
