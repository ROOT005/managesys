package controllers

import (
	"github.com/astaxie/beego"
	"managesys/models"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Time"] = time.Now()
	c.TplName = "index.tpl"
}
func (c *MainController) GetInfo() {
	//info := models.Getinfo()
	info := "Hello"
	beego.Error(models.GetInfo())
	c.Ctx.Output.JSON(info, true, true)
}
