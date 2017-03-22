package routers

import (
	"github.com/astaxie/beego"
	"github.com/qor/admin"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/validations"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/confirm"
	_ "gopkg.in/authboss.v0/recover"
	_ "gopkg.in/authboss.v0/register"
	"managesys/controllers"
	"managesys/db"
	"managesys/models"
	"net/http"
)

type AdminAuth struct{}

var (
	Auth = authboss.New()
)

func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/login"
}
func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/logout"
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	email, _ := c.Request.Cookie("id")
	value, err := c.Request.Cookie("see")
	if err == nil && value.Value == "BgQDwQ3THJn9F7NPLBi6hTI3Fwz55h47jQUVCOL6iq" {
		var user models.User
		if !db.DB.First(&user, "email = ?", email).RecordNotFound() {
			return &user
		}
	}
	return nil
}

func init() {
	//链接数据库
	DB := db.DB
	//本地化数据
	I18n := i18n.New(database.New(DB))
	//注册资源
	DB.AutoMigrate(&models.User{}, &models.Client{})
	Admin := admin.New(&qor.Config{DB: DB})

	/**************添加菜单***************/
	//控制面板
	Admin.AddMenu(&admin.Menu{Name: "Dashboard", Link: "/admin"})
	//管理员管理
	admin_auth := AdminAuth{}
	Admin.SetAuth(&admin_auth)
	adminuser := Admin.AddResource(&models.User{}, &admin.Config{Menu: []string{"User"}})
	adminuser.Meta(&admin.Meta{Name: "Password",
		Type:            "password",
		FormattedValuer: func(interface{}, *qor.Context) interface{} { return "" },
		Setter: func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			values := metaValue.Value.([]string)
			if len(values) > 0 {
				if newPassword := values[0]; newPassword != "" {
					bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
					if err != nil {
						context.DB.AddError(validations.NewError(adminuser, "Password", "无法加密密码"))
						return
					}
					u := resource.(*models.User)
					u.Password = string(bcryptPassword)
				}
			}
		},
	})
	adminuser.IndexAttrs("-Password")

	//客户数据
	client := Admin.AddResource(&models.Client{}, &admin.Config{Menu: []string{"Site Management"}})
	client.Meta(&admin.Meta{Name: "ClientInfo", Type: "single_edit"})
	client.Meta(&admin.Meta{Name: "AssetInfo", Type: "single_edit"})
	client.IndexAttrs("ID", "Name", "Count", "Level", "State", "Result")
	//添加翻译
	Admin.AddResource(I18n)

	/*****************路由************************************/
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	beego.Handler("/admin/*", mux)
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/", &controllers.MainController{})
}
