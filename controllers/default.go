package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "try ppp"
	c.Data["Email"] = "agaihghioooinjjjj@gmail.com"
	c.TplName = "index.tpl"
}
