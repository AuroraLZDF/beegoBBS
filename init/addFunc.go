package init

import (
	"auroraLZDF/beegoBBS/utils"
	"github.com/astaxie/beego"
)

func AddFunc() {
	beego.AddFuncMap("config", utils.Configs)
	beego.AddFuncMap("authCheck", utils.AuthCheck)
}
