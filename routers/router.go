package routers

import (
	"github.com/auroraLZDF/beegoBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
