package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "try again"
	c.Data["Email"] = "again@gmail.com"
	c.TplName = "index.tpl"
}
