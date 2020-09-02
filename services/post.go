package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pb/global"
	"pb/global/response"
	"pb/model"
)

func GetPostList(c *gin.Context) {
	var posts []model.Post
	var _posts []model.TransFormPost
	var topic model.Topic
	if err := global.GVA_DB.Find(&posts).Error; err != nil {
		response.FailWithMessage(fmt.Sprintf("发生了错误：&s", err), c)
	} else {
		for _, item := range posts {
			var topicIfExist bool
			if err := global.GVA_DB.First(&topic, item.TopicID).Error; err != nil {
				fmt.Println(topic, 1)
				topicIfExist = false
			} else {
				fmt.Println(topic, 2)
				topicIfExist = true
			}
			_posts = append(_posts, model.TransFormPost{ID: item.ID, Category: item.Category, TopicID: item.TopicID, TopicIfExist: topicIfExist, ShortMd: item.Md[:66]})
		}
		response.OkWithData(_posts, c)
	}

}

func CreatePost(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {

		response.FailWithMessage(fmt.Sprintf("发生了错误:%s", err), c)
	} else {
		if err := global.GVA_DB.Create(&post).Error; err != nil {
			response.FailWithMessage(fmt.Sprintf("failure %s", err), c)
		} else {
			response.OkWithData(post, c)
		}
	}

}

// 删除(。。。。。)
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post model.Post
	if err := global.GVA_DB.Unscoped().Delete(&post, id).Error; err != nil {
		response.FailWithMessage(fmt.Sprintf("删除失败:%s", err), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}

}

// 修改
