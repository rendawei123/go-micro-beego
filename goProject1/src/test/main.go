package main

import "github.com/astaxie/beego"

type HomeController struct {
	beego.Controller
}

func (self *HomeController) Get() {
	self.Ctx.WriteString("Hello Word")
}

func main() {
	beego.Router("/", &HomeController{})
	beego.Run()
}
