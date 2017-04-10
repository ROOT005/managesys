package controllers

import (
	//"encoding/json"
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

	weekInfo, weekInfoFin := models.GetWeekInfo()
	dayInfo, dayInfoFin := models.GetDayInfo()

	info := map[string]interface{}{
		"weekInfo":    weekInfo,
		"weekInfoFin": weekInfoFin,
		"dayInfo":     dayInfo,
		"dayInfoFin":  dayInfoFin,
	}
	c.Ctx.Output.JSON(info, true, true)
}
