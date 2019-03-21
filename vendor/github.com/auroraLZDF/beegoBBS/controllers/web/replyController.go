package web

import (
	"github.com/auroraLZDF/beegoBBS/models"
	"github.com/auroraLZDF/beegoBBS/utils"
)

type ReplyController struct {
	Controller
}

func (this *ReplyController) Store() {
	var topic_id = this.Ctx.Input.Param(":id")
	var body = this.GetString("body")

	var reply = models.Replies{
		TopicId: utils.StringToInt(topic_id),
		Content: body,
		UserId:  utils.FloatToInt(uInfo["id"].(float64)),
	}

	if err := reply.AddReply(reply); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	data["url"] = "/topics/show/" + topic_id
	this.JsonMessage(1, "评论添加成功", data)
}

func (this *ReplyController) Destroy() {
	var id = this.Ctx.Input.Param(":id")
	var topic_id = this.GetString("topic_id")

	if err := utils.Required(id); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Required(topic_id); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	reply := models.Replies{}
	if err := reply.Delete(utils.StringToInt(id), utils.StringToInt(topic_id)); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	data["url"] = "/topics/show/" + topic_id
	this.JsonMessage(1, "删除评论成功", data)
}
