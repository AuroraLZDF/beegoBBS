package upload

import "github.com/astaxie/beego"

type Controller struct {
	beego.Controller
}

var data = make(map[string]interface{})
var domain = beego.AppConfig.String("Url")
var uploadDir = beego.AppConfig.String("upload_dir")


func (this *Controller) JsonMessage(code int, msg string, data map[string]interface{}) {
	result := map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	this.Data["json"] = result
	this.ServeJSON()
}
