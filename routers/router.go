package routers

import (
	"github.com/auroraLZDF/beegoBBS/controllers"
	"github.com/astaxie/beego"
	"github.com/auroraLZDF/beegoBBS/controllers/auth"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	/** auth start **/
	// login
	beego.Router("/login", &auth.LoginController{}, "get:LoginForm")
	beego.Router("/login", &auth.LoginController{}, "post:Login")
	beego.Router("/logout", &auth.LoginController{}, "get:Logout")
	// register
	beego.Router("/register", &auth.RegisterController{}, "get:RegisterForm")
	beego.Router("/register", &auth.RegisterController{}, "post:Register")

	// reset password
	beego.Router("/password/forget", &auth.PasswordController{}, "get:Forget")
	beego.Router("/password/email", &auth.PasswordController{}, "post:SendResetLinkEmail")
	beego.Router("/password/reset/:token", &auth.PasswordController{}, "get:ShowResetForm")
	beego.Router("/password/reset", &auth.PasswordController{}, "post:Reset")

	/** auth end **/
}
