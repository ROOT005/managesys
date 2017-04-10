package routers

import (
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"github.com/qor/i18n"
	"github.com/qor/i18n/backends/database"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/roles"
	"github.com/qor/validations"
	"golang.org/x/crypto/bcrypt"
	//"gopkg.in/authboss.v0"
	_ "gopkg.in/authboss.v0/auth"
	_ "gopkg.in/authboss.v0/confirm"
	_ "gopkg.in/authboss.v0/recover"
	_ "gopkg.in/authboss.v0/register"
	"managesys/controllers"
	"managesys/db"
	"managesys/models"
	"net/http"
)

var Role string

type Auth struct{}

func (Auth) LoginURL(c *admin.Context) string {
	return "/login"
}
func (Auth) LogoutURL(c *admin.Context) string {
	return "/logout"
}

func (Auth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	email, _ := c.Request.Cookie("id")
	value, err := c.Request.Cookie("see")
	if err == nil && value.Value == "BgQDwQ3THJn9F7NPLBi6hTI3Fwz55h47jQUVCOL6iq" {
		var user models.User
		if !db.DB.First(&user, "email = ?", email.Value).RecordNotFound() {
			return &user
		}
	}
	return nil
}

func init() {
	/********************权限设置****************/
	roles.Register("超级管理员", func(req *http.Request, currentUser interface{}) bool {
		return req.RemoteAddr == "127.0.0.1" || (currentUser.(*models.User) != nil && currentUser.(*models.User).Role == "超级管理员")
	})
	roles.Register("店长", func(req *http.Request, currentUser interface{}) bool {
		return currentUser.(*models.User) != nil && currentUser.(*models.User).Role == "店长"
	})
	roles.Register("店员", func(req *http.Request, currentUser interface{}) bool {
		return currentUser.(*models.User) != nil && currentUser.(*models.User).Role == "店员"
	})

	//链接数据库
	DB := db.DB
	I18n := i18n.New(database.New(DB))
	DB.AutoMigrate(&models.User{}, &models.Client{}, &models.AssetInfo{}, &models.FullHouse{}, &models.MortgageHouse{}, &models.FullCar{}, &models.MortgageCar{}, &models.InsurancePolicy{}, &models.AccuFound{}, &models.SocialSecurity{}, &models.Salary{}, &models.BusinessLoan{})
	Admin := admin.New(&qor.Config{DB: DB})
	Admin.GetRouter().Get("/", func(c *admin.Context) {
		http.Redirect(c.Writer, c.Request, "/admin/clients?scopes=我的客户", http.StatusSeeOther)
	})

	/**************添加菜单***************/
	//管理员管理
	Admin.SetAuth(Auth{})
	Admin.SetSiteName("51DK管理系统")
	adminuser := Admin.AddResource(&models.User{}, &admin.Config{
		Menu:       []string{"User"},
		Permission: roles.Allow(roles.CRUD, "超级管理员").Allow(roles.CRUD, "店长").Deny(roles.CRUD, "店员"),
	})
	adminuser.Meta(&admin.Meta{Name: "Role", Config: &admin.SelectOneConfig{Collection: []string{"超级管理员", "店长", "店员"}}})
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
	client := Admin.AddResource(&models.Client{}, &admin.Config{
		Menu: []string{"Site Management"},
	})
	client.Meta(&admin.Meta{
		Name:       "Level",
		Config:     &admin.SelectOneConfig{Collection: []string{"优秀", "良好", "中等", "较差"}},
		Permission: roles.Allow(roles.CRUD, "超级管理员").Allow(roles.CRUD, "店长").Allow(roles.CRUD, "店员"),
	})
	client.Meta(&admin.Meta{
		Name:       "Result",
		Permission: roles.Allow(roles.CRUD, "超级管理员").Allow(roles.CRUD, "店长").Allow(roles.Read, "店员"),
	})
	client.Scope(&admin.Scope{Name: "我的客户", Handle: func(db *gorm.DB, context *qor.Context) *gorm.DB {
		roles := context.Roles[0]
		if roles == "超级管理员" || roles == "店长" {
			var clients []*models.Client
			return db.Find(&clients)
		} else {
			return db.Where("operator = ?", context.CurrentUser.DisplayName())
		}
	},
		Default: true,
	})
	client.Meta(&admin.Meta{
		Name:   "Gender",
		Config: &admin.SelectOneConfig{Collection: []string{"男", "女", "未知"}},
	})
	client.Meta(&admin.Meta{
		Name: "Operator",
		Setter: (func(resource interface{}, metaValue *resource.MetaValue, context *qor.Context) {
			operator := context.CurrentUser.DisplayName()
			u := resource.(*models.Client)
			u.Operator = operator
		}),
	})
	client.NewAttrs(
		&admin.Section{
			Title: "Basic InFo",
			Rows: [][]string{
				{"Operator"},
				{"Name", "Count", "Level"},
				{"State", "Result", "Other"},
				{"Gender", "Age", "PhoneNum"},
				{"Address", "MaritalStatuc", "Company"},
				{"CreditCheck", "CreCheckCount", "GetTime"},
				{"Balance", "CredCard", "PayAc"},
				{"MaxAc", "IdNum", "DetailedList"},
				{"CurrentAdd", "CompanyAdd"},
			},
		},
		&admin.Section{
			Title: "AssetInfo",
			Rows: [][]string{
				{"AssetInfo"},
			},
		},
	)
	client.EditAttrs(client.NewAttrs(), "-Operator")
	client.ShowAttrs("Operator", client.NewAttrs())

	client.IndexAttrs("ID", "Name", "Count", "Level", "State", "Result", "CreatedAt")

	client.Action(&admin.Action{
		Name: "批量删除",
		Handle: func(actionArgument *admin.ActionArgument) error {
			var client []*models.Client
			for _, record := range actionArgument.FindSelectedRecords() {
				actionArgument.Context.DB.Model(record.(*models.Client)).Delete(&client)
			}
			return nil
		},
		Modes: []string{"index"},
	})
	//添加翻译
	Admin.AddResource(I18n)

	/*****************路由************************************/
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	beego.Handler("/admin/*", mux)
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/info", &controllers.MainController{}, "get:GetInfo")
}
