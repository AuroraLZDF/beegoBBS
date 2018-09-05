package web

import (
	"github.com/auroraLZDF/beegoBBS/utils"
	"github.com/auroraLZDF/beegoBBS/models"
	"strconv"
	"fmt"
)

type UserController struct {
	Controller
}

func (this *UserController) user(id string) interface{} {
	_id, _ := strconv.Atoi(id)
	b, res, errStr := models.FindUserByFields(models.Users{Id: _id})
	if b == false {
		fmt.Println(errStr)
		return nil
	}

	return res
}

func (this *UserController) Show() {
	id := this.Ctx.Input.Param(":id")
	res := this.user(id)
	if res == nil {
		this.ShowError(res.(string))
		return
	}

	this.Data["user"] = utils.StructToMap(res)
	this.TplName = "web/user/index.html"
}

func (this *UserController) Edit() {
	id := this.Ctx.Input.Param(":id")
	res := this.user(id)
	if res == nil {
		this.ShowError(res.(string))
		return
	}

	this.Data["user"] = utils.StructToMap(res)
	this.TplName = "web/user/edit.html"
}

func (this *UserController) Update() {

}
