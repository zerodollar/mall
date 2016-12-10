package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Data["Website"] = "www.beego.me"
	c.Data["Email"] = "z29759@gmail.com"
	c.TplName = "index.tpl"
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.beego.me"
	c.Data["Email"] = "z29759@gmail.com"
	c.TplName = "index.tpl"
}
