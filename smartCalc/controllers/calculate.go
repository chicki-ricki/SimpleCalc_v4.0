package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type CalculateController struct {
	beego.Controller
}

func (c *CalculateController) Calculate() {
	c.TplName = "calculate/startCalculate.tpl"
}
