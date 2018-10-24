package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	//c.Ctx.WriteString(fmt.Sprint(c.Input()))  // 直接输出接收到的字符串

	// 获取用户名密码以及自动登录的值
	uname := c.Input().Get("uname")
	pwd := c.Input().Get("pwd")
	autoLogin := c.Input().Get("autoLogin") == "on"

	// 判断输入的是否是配置文件中的
	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		maxAge := 0 // 设置cookie保存时间为0
		if autoLogin {
			maxAge = 1<<31 - 1
		}

		// 设置cookie
		c.Ctx.SetCookie("uname", uname, maxAge, "/home")
		c.Ctx.SetCookie("pwd", pwd, maxAge, "/home")
	}
	// 跳转
	c.Redirect("/home", 301)
	return
}

// 检查cookie中的用户名密码和配置中的是否一致
func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname") // context.Context是啥？
	if err != nil {
		return false
	}
	uname := ck.Value

	ck1, err1 := ctx.Request.Cookie("pwd")
	if err1 != nil {
		return false
	}
	pwd := ck1.Value
	//fmt.Println(ck1.Value)
	//fmt.Println(beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd)
	return beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd
}
