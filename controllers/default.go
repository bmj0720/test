package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "tr22y1111www11222222111qq222pp"
	c.Data["Email"] = "agaihghioooinjjjj@gmail.com"
	c.TplName = "index.tpl"
}
