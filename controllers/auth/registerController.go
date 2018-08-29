package auth

import (
	"github.com/auroraLZDF/beegoBBS/utils"
	"github.com/auroraLZDF/beegoBBS/models"
	"github.com/auroraLZDF/beegoBBS/controllers"
)

type RegisterController struct {
	controllers.BaseController
}

func (this *RegisterController) RegisterForm() {
	this.TplName = "web/auth/register.html"
}

func (this *RegisterController) Register() {
	name := this.GetString("name")
	email := this.GetString("email")
	password := this.GetString("password")
	passwordConfirmation := this.GetString("password_confirmation")

	if _, err := utils.Required(name); err != nil {
		utils.ShowErr(err)
	}

	if _, err := utils.Required(password); err != nil {
		utils.ShowErr(err)
	}

	if _, err := utils.Required(passwordConfirmation); err != nil {
		utils.ShowErr(err)
	}

	if b, str := utils.Equal(password, passwordConfirmation); b == false {
		utils.ShowErr(str)
	}

	if _, err := utils.IsEmail(email); err != nil {
		utils.ShowErr(err)
	}


	if b, _, err := models.FindUserByFields(models.Users{Email: email}); b == true {
		utils.ShowErr(err)
	}

	user := models.Users{
		Name:          name,
		Email:         email,
		Password:      utils.Md5(password),
		RememberToken: utils.Md5(password + email),
	}
	if res, err := models.AddUser(user); res == false {
		utils.ShowErr(err)
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

	this.Redirect("/", 302)
}
