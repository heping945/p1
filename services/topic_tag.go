package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pb/global"
	"pb/global/response"
	"pb/model"
)

func GetTopicTagList(c *gin.Context) {
	var Taglist []model.Tag
	global.GVA_DB.Find(&Taglist)

	response.OkWithData(Taglist, c)
}

func CreateTopicTag(c *gin.Context) {
	var tag model.Tag
	if err := c.ShouldBind(&tag); err != nil {

		response.FailWithMessage(fmt.Sprintf("发生了错误:%s", err), c)
	} else {
		global.GVA_DB.Create(&tag)
		if tag.ID == 0 {
			response.FailWithMessage(fmt.Sprintf("failure 此分类已经创建过"), c)
		} else {
			response.OkWithData(tag, c)
		}
	}
}

func DeleteTopicTag() {

}
