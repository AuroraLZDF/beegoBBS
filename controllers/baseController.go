package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// run before get
func (this *BaseController) Prepare() {
	// login status
	/*login := this.GetSession("loginUser")
	if login == nil {
		this.Redirect("/login", 302)
	} else {
		this.Redirect("/", 200)
	}*/
	//this.SetSession("loginUser", "aurora")
	loginName := this.GetSession("loginUser")
	if loginName != nil && loginName != "" {
		this.Data["loginName"] = loginName
	}
}

