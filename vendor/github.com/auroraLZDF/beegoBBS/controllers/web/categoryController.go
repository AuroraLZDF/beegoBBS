package web

import (
	"github.com/auroraLZDF/beegoBBS/models"
	"github.com/auroraLZDF/beegoBBS/utils"
)

type CategoryController struct {
	Controller
}

//
func (this *CategoryController) Show() {
	id, _ := this.GetInt("id")
	page, _ := this.GetInt("page")
	category_id := this.Ctx.Input.Param(":id")

	if err := utils.Required(category_id); err != nil {
		this.Error404("不存在的话题分类")
		return
	}

	var order = "recent"
	where := map[string]interface{}{
		"id":          id,
		"category":    utils.StringToInt(category_id),
		"page":        page,
		"order":       order,
		"currentPath": utils.CurrentPath(this.Ctx.Request),
	}
	topic := models.Topics{}

	result, err := topic.TopicLists(where)
	if err != nil {
		this.Error404("话题列表加载失败")
	}

	this.Data["order"] = order
	this.Data["category_active"] = utils.StringToInt(category_id)
	this.Data["topics"] = result["topics"]
	this.Data["pageNate"] = result["pageNate"]
	this.TplName = "web/topic/index.html"
}
