package routers

import (
	"github.com/astaxie/beego"
	"auroraLZDF/beegoBBS/controllers/auth"
	"auroraLZDF/beegoBBS/controllers/upload"
	"auroraLZDF/beegoBBS/controllers/web"
	"github.com/dchest/captcha"
)

func init() {
	beego.Router("/", &web.TopicController{}, "get:Index")

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

	// user center
	beego.Router("/user/:id", &web.UserController{}, "get:Show")
	beego.Router("/user/edit/:id", &web.UserController{}, "get:Edit")
	beego.Router("/user/save", &web.UserController{}, "post:Update")

	// 上传文件
	beego.Router("/file/upload", &upload.WebUploadController{}, "post:Upload")

	// 验证码
	beego.Handler("/captcha/*.png", captcha.Server(240, 80))

	// 导航栏
	beego.Router("/categories/show/:id", &web.CategoryController{}, "get:Show")
	// Topic
	beego.Router("/topics", &web.TopicController{}, "get:Index")
	beego.Router("/topics/show/:id", &web.TopicController{}, "get:Show")
	beego.Router("/topics/create", &web.TopicController{}, "get:Create")
	beego.Router("/topics/store", &web.TopicController{}, "post:Save")
	beego.Router("/topics/edit/:id", &web.TopicController{}, "get:Edit")
	beego.Router("/topics/update", &web.TopicController{}, "post:Update")
	beego.Router("/topics/destroy", &web.TopicController{}, "post:Destroy")
	// 回复/留言
	beego.Router("/replies/store/:id", &web.ReplyController{}, "post:Store")
	beego.Router("/replies/destroy/:id", &web.ReplyController{}, "post:Destroy")
}
