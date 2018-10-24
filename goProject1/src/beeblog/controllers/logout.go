package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {

	//cookie, _ := c.Ctx.Request.Cookie("uname")
	//fmt.Println(cookie)

	c.Data["IsLogin"] = false
	c.Ctx.SetCookie("uname", "uuu", 0, "/home")
	c.Ctx.SetCookie("pwd", "uuu", 0, "/home")

	//cookie1, _ := c.Ctx.Request.Cookie("uname")
	//fmt.Println(cookie1)

	c.TplName = "home.html"
}
