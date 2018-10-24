package controllers

import (
	"beeblog/models"
	"fmt"
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")
	fmt.Println("。。。。。。。。。。。。", op)
	switch op {
	case "add":
		name := c.Input().Get("name")
		fmt.Println("------------", name)
		if len(name) == 0 {
			break
		}

		err := models.AddCategory(name)
		if err == nil {
			beego.Error(err)
		}
		c.Redirect("/home/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		fmt.Println(id)
		err1 := models.DeleteCategory(id)
		if err1 != nil {
			beego.Error(err1)
		}
		c.Redirect("/home/category", 301)
		return
	}
	c.TplName = "category.html"

	var err error
	c.Data["Categories"], err = models.GetAllCategories()

	if err != nil {
		beego.Error(err)
	}
}
