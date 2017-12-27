package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "ljchen.net"
	c.Data["Email"] = "chenleji@gmail.com"
	c.TplName = "index.tpl"
}