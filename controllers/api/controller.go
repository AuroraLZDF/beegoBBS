package api

import "auroraLZDF/beegoBBS/controllers"

type Controller struct {
	controllers.BaseController
}

var data = make(map[string]interface{})