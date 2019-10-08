// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"rbacAdmin/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	ns := beego.NewNamespace("/admin",

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/course",
			beego.NSInclude(
				&controllers.CourseController{},
			),
		),

		beego.NSNamespace("/video",
			beego.NSInclude(
				&controllers.VideoController{},
			),
		),

		beego.NSNamespace("/teacher",
			beego.NSInclude(
				&controllers.TeacherController{},
			),
		),

		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),

		beego.NSNamespace("/home",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),

		beego.NSNamespace("/rbac",
			beego.NSInclude(
				&controllers.RbacController{},
			),
		),

	)
	beego.AddNamespace(ns)

	beego.Get("/", func(c *context.Context) {

		//已登录，后台首页；未登录，登录页
		sess, _ := beego.GlobalSessions.SessionStart(c.ResponseWriter, c.Request)
		defer sess.SessionRelease(c.ResponseWriter)


		sessRole := sess.Get("role")
		if sessRole != nil {
			c.Redirect(302, "/admin/user")
		}else{
			c.Redirect(302, "/admin/login/login")
		}
	})

	//beego.Router("/home", &controllers.MainController{})
}
