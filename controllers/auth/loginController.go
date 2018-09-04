package auth

import (
	"github.com/auroraLZDF/beegoBBS/models"
	"github.com/auroraLZDF/beegoBBS/utils"
)

type LoginController struct {
	Controller
}

func (this *LoginController) LoginForm() {
	this.TplName = "web/auth/login.html"
}

func (this *LoginController) Login() {
	email := this.Input().Get("email")
	password := this.Input().Get("password")

	if b, err := utils.IsEmail(email); b == false {
		this.JsonMessage(2, err.Error(), data)
	}

	if b, err := utils.Required(password); b == false {
		this.JsonMessage(2, err.Error(), data)
	}

	b, user, err := models.FindUserByFields(models.Users{Email:email,Password:utils.Md5(password)})
	if b == false {
		this.JsonMessage(2, err, data)
	}

	uInfo := map[string]interface{}{
		"id": user.Id,
		"name": user.Name,
		"email": user.Email,
		"password": user.Password,
	}
	js := utils.MapToJson(uInfo)

	this.SetSession("uInfo", utils.AuthCode(js, "encode"))

	data["url"] = "/"
	this.JsonMessage(1, "登录成功！", data)
}

func (this *LoginController) Logout() {
	this.DelSession("uInfo")
	this.Redirect("/login", 302)
}
