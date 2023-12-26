package main

import (
	d "smartCalc/domains"
	_ "smartCalc/routers"
	t "smartCalc/tools"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	t.FileCheck(d.NeccessoryFiles)
	beego.BConfig.Listen.Graceful = true

	defer t.Clg.LogFile.Close()

	beego.Run()

}
