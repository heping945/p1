package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pb/global"
	"pb/global/response"
	"pb/model"
)

// 创建一个主题
func CreateTopic(c *gin.Context) {
	var topic model.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {

		response.FailWithMessage(fmt.Sprintf("发生了错误:%s", err), c)
	} else {
		if err := global.GVA_DB.Create(&topic).Error; err != nil {
			response.FailWithMessage(fmt.Sprintf("failure %s", err), c)
		} else {
			response.OkWithData(topic, c)
		}
	}
}

func GetTopicList(c *gin.Context) {
	var topic []model.Topic
	var _topic []model.TransFormTopic
	global.GVA_DB.Find(&topic)

	for _, item := range topic {
		_topic = append(_topic, model.TransFormTopic{ID: item.ID, Title: item.Title, Difficulty: item.Difficulty, Slug: item.Slug})
	}

	response.OkWithData(_topic, c)
}

// 获得单个主题
func GetTopic(c *gin.Context) {
	slug := c.Param("slug")
	var topic model.Topic

	if err := global.GVA_DB.Preload("Post").Where(map[string]interface{}{"slug": slug}).First(&topic).Error; err != nil {
		response.FailWithMessage(fmt.Sprintf("未获取到参数为 %s 数据 ", slug), c)
	} else {
		response.OkWithData(topic, c)
	}
}

// 删除
func DeleteTopic(c *gin.Context) {
	slug := c.Param("slug")
	var topic []model.Topic
	global.GVA_DB.Where("slug = ?", slug).Unscoped().Delete(&topic)

	response.OkWithMessage("删除成功", c)

}

// 修改
func UpdateTopic() {

}
