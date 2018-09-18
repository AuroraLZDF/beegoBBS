package web

import (
	"html/template"

	"github.com/auroraLZDF/beegoBBS/controllers"
	"github.com/auroraLZDF/beegoBBS/utils"
)

type Controller struct {
	controllers.BaseController
}

var data = make(map[string]interface{})

var uInfo = make(map[string]interface{})

// run before get
func (this *Controller) Prepare() {
	// XSRF
	this.XSRFExpire = 7200
	this.Data["xsrf_html"] = template.HTML(this.XSRFFormHTML())
	this.Data["xsrf_token"] = this.XSRFToken()
	this.Data["category_active"] = 0

	if res := this.CheckCk(); res != nil {
		uInfo = res
		this.Data["uInfo"] = res
		return
	}

	this.Redirect("/login", 302)
}

func (this *Controller) CheckMe(id int) error {

	if id != utils.FloatToInt(uInfo["id"].(float64)) {
		return utils.Error("对不起，您无权访问！")
	}

	return nil
}
