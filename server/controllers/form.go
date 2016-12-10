package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type user struct {
	Id    int         `form:"-"`
	Name  interface{} `orm:"size(100)" form:"username" valid:"Required;Match(/^Bee.*/)"`
	Age   int         `                form:"age"      valid:"Range(1, 140)"`
	Email string      `orm:"size(100)"                 valid:"Email; MaxSize(100)"`
}

func (u *user) Valid(v *validation.Validation) {
}

type DemoController struct {
	beego.Controller
}

func (c *DemoController) Post() {
	u := user{}
	if err := c.ParseForm(&u); err != nil {
		//handle error
	}
}
