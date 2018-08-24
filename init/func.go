package init

import (
	"github.com/auroraLZDF/beegoBBS/utils"
	"github.com/astaxie/beego"
)

func AddFunc() {
	beego.AddFuncMap("config", utils.Configs)
	beego.AddFuncMap("checkCk", utils.CheckCk)
}
