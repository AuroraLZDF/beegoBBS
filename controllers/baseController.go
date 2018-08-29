package controllers

import (
	"github.com/astaxie/beego"
	"github.com/auroraLZDF/beegoBBS/utils"
	"github.com/auroraLZDF/beegoBBS/models"
)

type BaseController struct {
	beego.Controller
}

// run before get
func (this *BaseController) Prepare() {

	if res := this.CheckCk();res != nil {
		this.Data["uInfo"] = res
	}
}

func (this *BaseController) CheckCk() map[string]interface{} {
	uInfo := this.GetSession("uInfo")

	if uInfo != nil && uInfo != "" {
		js := utils.AuthCode(uInfo.(string), "DECODE")
		m := utils.JsonToMap(js)

		// Golang 使用 JSON unmarshal 数字到 interface{} 数字变成 float64 类型
		// 将 “id” 键申明为 float64 类型，再转换为 int 型
		id := int(m["id"].(float64))

		user := models.Users{
			//Id: m["id"].(int),
			Id:  id,
			Name:m["name"].(string),
			Email:m["email"].(string),
			Password:m["password"].(string),
		}
		if b, _, err := models.FindUserByFields(user); b == false {
			utils.ShowErr(err)
		}

		return m
	}

	return nil
}
