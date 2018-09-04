package auth

import (
	"github.com/auroraLZDF/beegoBBS/utils"
	"github.com/auroraLZDF/beegoBBS/models"
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

	if _, err := utils.Required(name); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if _, err := utils.Required(password); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if _, err := utils.Required(passwordConfirmation); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	if b, str := utils.Equal(password, passwordConfirmation); b == false {
		this.JsonMessage(2, str, data)
		return
	}

	if _, err := utils.IsEmail(email); err != nil {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	captchaId := this.GetString("captchaId")
	captchaValue := this.GetString("captcha")
	if !captcha.VerifyString(captchaId, captchaValue) {
		this.JsonMessage(2, "请填写正确的验证码", data)
		return
	}

	if b, _, err := models.FindUserByFields(models.Users{Email: email}); b == true {
		this.JsonMessage(2, err, data)
		return
	}

	user := models.Users{
		Name:          name,
		Email:         email,
		Password:      utils.Md5(password),
		RememberToken: utils.Md5(password + email),
	}
	if res, err := models.AddUser(user); res == false {
		this.JsonMessage(2, err.Error(), data)
		return
	}

	_, res, _ := models.FindUserByFields(user)

	uInfo := map[string]interface{}{
		"id": res.Id,
		"name": res.Name,
		"email": res.Email,
		"password": res.Password,
	}
	js := utils.MapToJson(uInfo)

	this.SetSession("uInfo", utils.AuthCode(js, "encode"))

	data["url"] = "/"
	this.JsonMessage(1, "注册成功！", data)
	return
}
