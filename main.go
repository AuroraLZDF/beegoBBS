package main

import (
	_ "github.com/auroraLZDF/beegoBBS/routers"
	_ "github.com/auroraLZDF/beegoBBS/init"
	"github.com/astaxie/beego"
)

/*func init() {

}*/

func main() {
	beego.Run()
}
