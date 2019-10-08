package controllers

import (
	"github.com/astaxie/beego"
	"rbacAdmin/repositories"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

//@router /welcome	[get]
func (c *MainController) Welcome() {

	c.Layout = "_layout/layout.html"
	c.TplName = "home/welcome.html"
	c.LayoutSections = map[string]string{}
	c.LayoutSections["HtmlHead"] = "_layout/header.html"
	c.LayoutSections["Scripts"] = "_layout/scripts.html"
	c.LayoutSections["HtmlFooter"] = "_layout/footer.html"
	c.LayoutSections["LeftNav"] = "_layout/left_nav.html"
	c.Data["weburl"] = beego.AppConfig.String("weburl")

	c.Data["menus"] = c.GetSession("menus").(map[int]*repositories.Menus)

	//当前登录用户信息
	c.Data["admin"] = c.GetSession("user")
}
