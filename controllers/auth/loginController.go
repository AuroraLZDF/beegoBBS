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

	if err := utils.IsEmail(email); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Required(password); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	user, err := models.FindUserByFields(models.Users{Email:email,Password:utils.Md5(password)})
	if err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	uInfo := map[string]interface{}{
		"id": user.Id,
		"name": user.Name,
		"email": user.Email,
		"password": user.Password,
		"avatar": user.Avatar,
	}
	js := utils.MapToJson(uInfo)

	this.SetSession("uInfo", utils.AuthCode(js, "encode"))

	data["url"] = "/"
	this.JsonMessage(1, "登录成功！", data)
	return
}

func (this *LoginController) Logout() {
	this.DelSession("uInfo")
	this.Redirect("/login", 302)
}
