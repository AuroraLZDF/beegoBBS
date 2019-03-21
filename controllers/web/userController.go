package web

import (
	"log"
	"time"

	"auroraLZDF/beegoBBS/models"
	"auroraLZDF/beegoBBS/utils"
)

type UserController struct {
	Controller
}

func (this *UserController) user(id string) interface{} {
	_id := utils.StringToInt(id)
	res, err := models.FindUserByFields(models.Users{Id: _id})
	if err != nil {
		log.Fatal("ERROR | ", time.Now().Format("2006-01-02 15:04:05"), " | ", err)
		return nil
	}

	return res
}

func (this *UserController) Show() {
	id := this.Ctx.Input.Param(":id")
	res := this.user(id)
	if res == nil {
		this.Error404(res.(string))
		return
	}

	this.Data["user"] = utils.StructToMap(res)
	this.TplName = "web/user/index.html"
}

func (this *UserController) Edit() {
	id := this.Ctx.Input.Param(":id")

	// 检测是否当前用户
	if err := this.CheckMe(utils.StringToInt(id)); err != nil {
		this.Error403(err.Error())
		return
	}

	res := this.user(id)
	if res == nil {
		this.Error404(res.(string))
		return
	}

	this.Data["user"] = utils.StructToMap(res)
	this.TplName = "web/user/edit.html"
}

func (this *UserController) Update() {
	id, _ := this.GetInt("id")

	// 检测是否当前用户
	if err := this.CheckMe(id); err != nil {
		this.Error403(err.Error())
		return
	}

	name := this.GetString("name")
	email := this.GetString("email")
	introduction := this.GetString("introduction")
	avatar := this.GetString("avatar")

	if err := utils.Numeric(id); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Required(name); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.IsEmail(email); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Required(introduction); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Required(avatar); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	user := models.Users{
		Name:         name,
		Email:        email,
		Introduction: introduction,
		Avatar:       avatar,
	}
	if err := models.UpdateUserById(id, user); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	data["url"] = "/user/" + utils.IntToString(id)
	this.JsonMessage(1, "修改个人信息成功", data)
}
