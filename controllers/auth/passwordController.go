package auth

import (
	"github.com/auroraLZDF/beegoBBS/models"
	"github.com/auroraLZDF/beegoBBS/utils"
	"github.com/astaxie/beego"
)

type PasswordController struct {
	Controller
}

func (this *PasswordController) Forget() {
	this.TplName = "web/auth/password/forget.html"
}

func (this *PasswordController) SendResetLinkEmail() {

	email := this.GetString("email")
	if err := utils.IsEmail(email); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	res, err := models.FindUserByFields(models.Users{Email: email})
	if err != nil {
		this.JsonMessage(2, "该邮箱未被注册", data)
		return
	}

	// 发送邮件
	token := utils.Md5(res.Email + res.Password)

	if err := models.UpdateUserByEmail(email, models.Users{RememberToken: utils.TrimS(token)}); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	url := beego.AppConfig.String("Url")
	body := "请点击<a href='" + url + "/password/reset/" + token + "'>【重置密码】</a>该链接，重置密码。链接有效期为 10 分钟"
	if err := utils.SendMail(email, "重置密码", body, "html"); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	// 邮件发送成功，通知前台
	msg := "重置密码邮件发送成功，请在 10 分钟内完成修改密码操作！"
	data["token"] = token
	data["url"] = "/"
	this.JsonMessage(1, msg, data)
	return
}

func (this *PasswordController) ShowResetForm() {
	token := this.Ctx.Input.Param(":token")
	if err := utils.Required(token); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	user, err := models.FindUserByFields(models.Users{RememberToken: token})
	if err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	email := user.Email

	this.Data["email"] = email
	this.Data["token"] = token
	this.TplName = "web/auth/password/reset.html"
}

func (this *PasswordController) Reset() {
	token := this.GetString("token")
	email := this.GetString("email")
	password := this.GetString("password")
	password_confirmation := this.GetString("password_confirmation")

	if err := utils.IsEmail(email); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Required(token); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Equal(password, password_confirmation); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := models.UpdateUserByEmail(email, models.Users{RememberToken: "", Password: utils.Md5(password)}); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	data["url"] = "/"
	this.JsonMessage(1, "修改密码成功！", data)
	return
}
