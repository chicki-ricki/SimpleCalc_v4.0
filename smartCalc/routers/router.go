package routers

import (
	"smartCalc/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/calculate", &controllers.CalculateController{}, "get:Calculate")
	beego.Router("/calculate/start", &controllers.CalculateController{}, "*:Start")
}
