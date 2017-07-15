package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "tr22ywwww3333333ww122233311 pqqq222pp"
	c.Data["Email"] = "agaihghioooinjjjj@gmail.com"
	c.TplName = "index.tpl"
}
