package main

import (
	beego "github.com/beego/beego/v2/server/web"
	d "smartCalc/domains"
	_ "smartCalc/routers"
	t "smartCalc/tools"
)

func main() {

	t.FileCheck(d.NeccessoryFiles)
	beego.BConfig.Listen.Graceful = false

	defer t.Clg.LogFile.Close()

	beego.Run()

}
