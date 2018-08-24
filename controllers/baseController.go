package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

// run before get
/*func (this *BaseController) Prepare() {
	// login status
	user := this.GetSession("username")
	if user == nil {
		this.Redirect("/login", 302)
	} else {
		// find user id
		username := user.(string)
		u, err := models.FindUser(username)
		if err != nil {
			log.Fatal(err)
		}
	}


}*/

