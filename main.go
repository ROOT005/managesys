package main

import (
	"github.com/astaxie/beego"
	_ "managesys/routers"
)

func main() {
	beego.SetStaticPath("/system", "public/system")
	beego.SetStaticPath("/admin/asset", "public/admin/asset")
	beego.Run()
}
