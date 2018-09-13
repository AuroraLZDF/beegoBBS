package web

import (
	"github.com/auroraLZDF/beegoBBS/models"
	"github.com/auroraLZDF/beegoBBS/utils"
)

type TopicController struct {
	Controller
}

func (this *TopicController) Index() {
	id, _ := this.GetInt("id")
	category, _ := this.GetInt("category")
	page, _ := this.GetInt("page")
	order := this.Ctx.Input.Param(":order")

	where := map[string]interface{}{
		"id":       id,
		"category": category,
		"order":    order,
		"page":     page,
	}
	topic := models.Topics{}
	result, err := topic.TopicLists(where)
	if err != nil {
		this.Error404("话题列表加载失败")
	}

	//fmt.Println(result)
	//return

	this.Data["order"] = order
	this.Data["topics"] = result
	this.TplName = "web/topic/index.html"
}

func (this *TopicController) Show() {
	id := this.Ctx.Input.Param(":id")
	if err := utils.Required(id); err != nil {
		this.Error404(err.Error())
	}

	result, err := models.Topics{}.TopicByID(utils.StringToInt(id))
	if err != nil {
		this.Error404(err.Error())
	}

	this.Data["topic"] = result
	this.TplName = "web/topic/show.html"
}

func (this *TopicController) Create() {
	categories, err := models.Categories{}.FindAll()
	if err != nil {
		this.Error404(err.Error())
	}

	this.Data["categories"] = categories
	this.TplName = "web/topic/create.html"
}

func (this *TopicController) Save() {
	userId, _ := this.GetInt("user_id")
	categoryId, _ := this.GetInt("category_id")
	title := this.GetString("title")
	body := this.GetString("body")

	if err := utils.Numeric(userId); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Numeric(categoryId); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Required(title); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Required(body); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	topic := models.Topics{
		Title:      title,
		CategoryId: categoryId,
		Body:       body,
		UserId:     userId,
	}
	if err := topic.Create(topic); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	data["url"] = "/topics/recent"
	this.JsonMessage(1, "话题创建成功", data)
}

func (this *TopicController) Edit() {

}

func (this *TopicController) Update() {

}

func (this *TopicController) Destroy() {

}
