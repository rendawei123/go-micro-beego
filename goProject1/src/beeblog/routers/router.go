package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.HomeController{})
	beego.Router("/home/login", &controllers.LoginController{})
	beego.Router("/home/logout", &controllers.LogoutController{})
	beego.Router("/home/category", &controllers.CategoryController{})
}
