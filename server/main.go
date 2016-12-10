package main

import (
	"github.com/astaxie/beego"
	"github.com/zerodollar/mall/server/controllers"
	_ "github.com/zerodollar/mall/server/routers"
)

func main() {
	beego.SetLogger("file", `{"filename":"logs/test.log"}`)
	beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
