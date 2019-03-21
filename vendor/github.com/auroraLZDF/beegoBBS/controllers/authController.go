package controllers

type AuthController struct {
	BaseController
}

// run before get
func (this *AuthController) Prepare() {

	if res := this.CheckCk(); res != nil {
		this.Data["uInfo"] = res
		return
	}

	this.Redirect("/login", 302)
}
