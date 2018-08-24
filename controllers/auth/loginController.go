package auth

import (
	"github.com/auroraLZDF/beegoBBS/controllers"
	"log"
	"github.com/auroraLZDF/beegoBBS/models"
	"fmt"
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

	//fmt.Println(email, password)

	user := models.FindUserByEmail(email)
	/*uid := user.Id
	if (int(uid) != 0) {
		log.Fatal("err")
	}*/
	//os.Exit(1)
	fmt.Println(user)
}

func (this *LoginController) Logout() {

}
