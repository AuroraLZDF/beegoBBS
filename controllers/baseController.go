package controllers

import (
	"github.com/astaxie/beego"
	"github.com/auroraLZDF/beegoBBS/utils"
	"github.com/auroraLZDF/beegoBBS/models"
	"html/template"
	"log"
	"time"
)

type BaseController struct {
	beego.Controller
}

// run before get
func (this *BaseController) Prepare() {
	// XSRF
	this.XSRFExpire = 7200
	this.Data["xsrf_html"] = template.HTML(this.XSRFFormHTML())
	this.Data["xsrf_token"] = this.XSRFToken()

	if res := this.CheckCk(); res != nil {
		this.Data["uInfo"] = res
	}
}

func (this *BaseController) CheckCk() map[string]interface{} {
	uInfo := this.GetSession("uInfo")

	if uInfo != nil && uInfo != "" {
		js := utils.AuthCode(uInfo.(string), "DECODE")
		uMap := utils.JsonToMap(js)

		// Golang 使用 JSON unmarshal 数字到 interface{} 数字变成 float64 类型
		// 将 “id” 键申明为 float64 类型，再转换为 int 型
		id := int(uMap["id"].(float64))

		user := models.Users{
			//Id: m["id"].(int),
			Id:       id,
			Name:     uMap["name"].(string),
			Email:    uMap["email"].(string),
			Password: uMap["password"].(string),
		}
		if b, _, err := models.FindUserByFields(user); b == false {
			log.Fatal("ERROR | ", time.Now().Format("2006-01-02 15:04:05"), " | ", err)
			return nil
		}

		return uMap
	}

	return nil
}

func (this *BaseController) JsonMessage(code int, msg string, data map[string]interface{}) {
	result := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	this.Data["json"] = result
	this.ServeJSON()
}

func (this *BaseController) ShowError(err string) {
	this.Ctx.WriteString(err)
}
