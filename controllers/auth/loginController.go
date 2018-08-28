package auth

import (
	"github.com/auroraLZDF/beegoBBS/controllers"
	"log"
	"github.com/auroraLZDF/beegoBBS/models"
	"github.com/auroraLZDF/beegoBBS/utils"
)

type LoginController struct {
	controllers.BaseController
}

func (this *LoginController) LoginForm() {
	this.TplName = "web/auth/login.html"
}

func (this *LoginController) Login() {
	email := this.Input().Get("email")
	password := this.Input().Get("password")
	if email == "" {
		log.Fatal("邮箱不能为空")
	}
	if password == "" {
		log.Fatal("密码不能为空")
	}

	b, user, str := models.FindUserByFields(models.Users{Email:email,Password:utils.Md5(password)})
	if b == false {
		utils.ShowErr(str)
	}

	this.SetSession("loginUser", user.Email)

	this.Redirect("/", 302)
}

func (this *LoginController) Logout() {
	this.DelSession("loginUser")
	this.Redirect("/login", 302)
}
