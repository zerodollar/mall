package routers

import (
	"github.com/astaxie/beego"
	"github.com/zerodollar/mall/server/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
