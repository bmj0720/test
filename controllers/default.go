package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "try111 pqqq222pp"
	c.Data["Email"] = "agaihghioooinjjjj@gmail.com"
	c.TplName = "index.tpl"
}
