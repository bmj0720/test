package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "tr22222pp"
	c.Data["Email"] = "agaihgsshioooinjjjj@gmail.com"
	c.TplName = "index.tpl"
}
