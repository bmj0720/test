package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "tr2222111pp"
	c.Data["Email"] = "agaihgss111hioooinjjjj@gmail.com"
	c.TplName = "index.tpl"
}
