package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "smartCalc.web"
	c.Data["Email"] = "mlarra@student.21-school.ru"
	c.TplName = "index.tpl"
}
