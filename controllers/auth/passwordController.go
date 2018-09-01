package auth

import (
	"github.com/auroraLZDF/beegoBBS/controllers"
	"github.com/auroraLZDF/beegoBBS/models"
	"github.com/auroraLZDF/beegoBBS/utils"
	"github.com/astaxie/beego"
)

type PasswordController struct {
	controllers.BaseController
}

var data = make(map[string]interface{})

func (this *PasswordController) ShowLinkRequestForm() {
	this.TplName = "web/auth/password/reset.html"
}

func (this *PasswordController) Forget() {
	this.TplName = "web/auth/password/forget.html"
}

func (this *PasswordController) SendResetLinkEmail() {



	email := this.GetString("email")

	b, res, _ := models.FindUserByFields(models.Users{Email:email})
	if b == false {
		this.JsonMessage(2, "该邮箱未被注册", data)
	}

	// 发送邮件
	token := utils.Md5(res.Email + res.Password)
	url := beego.AppConfig.String("Url")
	body := "请点击<a href='" + url + "/password/reset/" + token + "'>【重置密码】</a>该链接，重置密码。链接有效期为 10 分钟"
	if err := utils.SendMail(email, "重置密码",body,"html"); err != nil {
		this.JsonMessage(2, err.Error(), data)
	}

	// 邮件发送成功，通知前台
	msg := "重置密码邮件发送成功，请在 10 分钟内完成修改密码操作！"
	data = map[string]interface{}{
		"token": token,
		"url": "/",
	}
	this.JsonMessage(1, msg, data)
}

func (this *PasswordController) ShowResetForm() {
	token := this.Ctx.Input.Param(":token")

	this.JsonMessage(1, token, data)
}

func (this *PasswordController) Reset() {

}
