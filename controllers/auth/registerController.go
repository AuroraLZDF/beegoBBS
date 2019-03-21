package auth

import (
	"auroraLZDF/beegoBBS/utils"
	"auroraLZDF/beegoBBS/models"
	"github.com/dchest/captcha"
)

type RegisterController struct {
	Controller
}

func (this *RegisterController) RegisterForm() {
	CaptchaId := captcha.NewLen(6)

	this.Data["CaptchaId"] = CaptchaId
	this.TplName = "web/auth/register.html"
}

func (this *RegisterController) Register() {
	name := this.GetString("name")
	email := this.GetString("email")
	password := this.GetString("password")
	passwordConfirmation := this.GetString("password_confirmation")

	if err := utils.Required(name); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.Equal(password, passwordConfirmation); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if err := utils.IsEmail(email); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	captchaId := this.GetString("captchaId")
	captchaValue := this.GetString("captcha")
	if !captcha.VerifyString(captchaId, captchaValue) {
		this.JsonMessage(2, "请填写正确的验证码", data)
		return
	}

	if _, err := models.FindUserByFields(models.Users{Email: email}); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	user := models.Users{
		Name:          name,
		Email:         email,
		Password:      utils.Md5(password),
		RememberToken: utils.Md5(password + email),
	}
	if err := models.AddUser(user); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	res, err := models.FindUserByFields(user)
	if err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	uInfo := map[string]interface{}{
		"id": res.Id,
		"name": res.Name,
		"email": res.Email,
		"password": res.Password,
		"avatar": user.Avatar,
	}
	js := utils.MapToJson(uInfo)

	this.SetSession("uInfo", utils.AuthCode(js, "encode"))

	data["url"] = "/"
	this.JsonMessage(1, "注册成功！", data)
	return
}
