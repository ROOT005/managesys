package controllers

import (
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"managesys/db"
	"managesys/models"
)

//登录
type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}
func (c *LoginController) Post() {
	email := c.Input().Get("email")
	value := "BgQDwQ3THJn9F7NPLBi6hTI3Fwz55h47jQUVCOL6iq"
	password := c.Input().Get("password")
	var user models.User
	if !db.DB.First(&user, "email = ?", email).RecordNotFound() {
		var passwd string
		passwd = user.Password
		err := bcrypt.CompareHashAndPassword([]byte(passwd), []byte(password))
		if err == nil {
			c.Ctx.SetCookie("id", email, "/")
			c.Ctx.SetCookie("see", value, "/")
			c.Redirect("/admin", 301)
			return
		} else {
			c.Redirect("/", 301)
		}
	} else {
		c.Redirect("/", 301)
	}

}

//退出

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {
	beego.Error("sahodhawoua\n")
	c.Ctx.SetCookie("id", "", -1, "/")
	c.Ctx.SetCookie("see", "", -1, "/")
	c.Redirect("/login", 301)
}
