package auth

import "github.com/auroraLZDF/beegoBBS/controllers"

type RegisterController struct {
	controllers.BaseController
}

func (this *RegisterController) RegisterForm() {
	this.TplName = "web/auth/register.html"
}

func (this *RegisterController) Register() {

}
