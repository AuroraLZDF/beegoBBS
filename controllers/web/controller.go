package web

import "github.com/auroraLZDF/beegoBBS/controllers"

type Controller struct {
	controllers.BaseController
}

var data = make(map[string]interface{})

// run before get
func (this *Controller) Prepare() {

	if res := this.CheckCk(); res != nil {
		this.Data["uInfo"] = res
		return
	}

	this.Redirect("/login", 302)
}
