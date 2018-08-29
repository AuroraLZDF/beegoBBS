package auth

import (
	"log"
	"github.com/auroraLZDF/beegoBBS/models"
	"github.com/auroraLZDF/beegoBBS/utils"
	"github.com/auroraLZDF/beegoBBS/controllers"
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

	uInfo := map[string]interface{}{
		"id": user.Id,
		"name": user.Name,
		"email": user.Email,
		"password": user.Password,
	}
	js := utils.MapToJson(uInfo)

	this.SetSession("uInfo", utils.AuthCode(js, "encode"))
	this.Redirect("/", 302)
}

func (this *LoginController) Logout() {
	this.DelSession("uInfo")
	this.Redirect("/login", 302)
}
