package web

import (
	"fmt"

	"auroraLZDF/beegoBBS/models"
	"auroraLZDF/beegoBBS/utils"
)

type TopicController struct {
	Controller
}

func (this *TopicController) Index() {
	id, _ := this.GetInt("id")
	categoryID, _ := this.GetInt("category_id")
	page, _ := this.GetInt("page")
	order := this.GetString("order")

	where := map[string]interface{}{
		"id":          id,
		"category":    categoryID,
		"order":       order,
		"page":        page,
		"currentPath": utils.CurrentPath(this.Ctx.Request),
	}
	topic := models.Topics{}
	result, err := topic.TopicLists(where)
	if err != nil {
		this.Error404("话题列表加载失败")
	}

	this.Data["order"] = order
	this.Data["category_active"] = categoryID
	this.Data["topics"] = result["topics"]
	this.Data["pageNate"] = result["pageNate"]
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

	// 判断话题作者是否是当前登录用户
	var me = true
	if err := this.CheckMe(result.UserId); err != nil {
		me = false
	}

	// 判断当前登录用户是否有删除评论权限
	/*
		var isAuthor = make(map[int]bool)
		for i, reply := range result.Reply {
			var ok = true
			if err := this.CheckMe(reply.UserId); err != nil {
				ok = false
			}
			isAuthor[i] = ok
		}
		this.Data["isAuthor"] = isAuthor
	*/

	this.Data["topic"] = result
	this.Data["CheckMe"] = me
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

	if userId != utils.FloatToInt(uInfo["id"].(float64)) {
		this.JsonMessage(2, "请使用当前用户创建话题", data)
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

	data["url"] = "/topics?order=recent"
	this.JsonMessage(1, "话题创建成功", data)
}

func (this *TopicController) Edit() {
	id := this.Ctx.Input.Param(":id")
	if err := utils.Required(id); err != nil {
		this.Error404(err.Error())
	}

	result, err := models.Topics{}.TopicByID(utils.StringToInt(id))
	if err != nil {
		this.Error404(err.Error())
	}

	if err := this.CheckMe(result.UserId); err != nil {
		this.Error403(err.Error())
	}

	categories, err := models.Categories{}.FindAll()
	if err != nil {
		this.Error404(err.Error())
	}

	this.Data["categories"] = categories
	this.Data["topic"] = result
	this.TplName = "web/topic/edit.html"
}

func (this *TopicController) Update() {
	id, _ := this.GetInt("id")
	userId, _ := this.GetInt("user_id")
	categoryId, _ := this.GetInt("category_id")
	title := this.GetString("title")
	body := this.GetString("body")

	if err := utils.Numeric(id); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

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

	if userId != utils.FloatToInt(uInfo["id"].(float64)) {
		fmt.Println(userId, utils.FloatToInt(uInfo["id"].(float64)))
		this.JsonMessage(2, "请使用当前用户创建话题", data)
		return
	}

	topic := models.Topics{
		Title:      title,
		CategoryId: categoryId,
		Body:       body,
		UserId:     userId,
	}

	if err := topic.UpdateById(id, topic); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	data["url"] = "/topics"
	this.JsonMessage(1, "话题编辑成功", data)
}

func (this *TopicController) Destroy() {
	id, _ := this.GetInt("id")
	userId, _ := this.GetInt("user_id")

	if err := utils.Numeric(id); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Numeric(userId); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if userId != utils.FloatToInt(uInfo["id"].(float64)) {
		fmt.Println(userId, utils.FloatToInt(uInfo["id"].(float64)))
		this.JsonMessage(2, "对不起，您没有删除该话题的权限", data)
		return
	}

	var topic = models.Topics{}
	if err := topic.DeleteById(id); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	data["url"] = "/topics"
	this.JsonMessage(1, "删除话题成功", data)
}
