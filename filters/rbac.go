package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"rbacAdmin/models"
)

/**
 *@author LanguageY++2013
 *2019/2/27 7:05 PM
 **/
func init() {
	beego.InsertFilter("/admin/*", beego.BeforeRouter, func(context *context.Context) {
		//校验用户名和密码
		r, w := context.Request, context.ResponseWriter
		sess, _ := beego.GlobalSessions.SessionStart(w, r)
		defer sess.SessionRelease(w)

		//登录无需验证授权
		if r.URL.Path == "/admin/login/login"{
			return
		}

		//获取权限列表
		canAccess := false
		spl := sess.Get("permissions")
		if spl == nil {//session过期
			context.Redirect(302, "/admin/login/login")
			return
		}

		if r.URL.Path == "/admin/home/welcome" || r.URL.Path == "/admin/login/login_out" {
			return
		}


		pl := spl.([]models.AdminPermissions)
		for _, p := range pl {
			if p.UrlPath == r.URL.Path {
				canAccess = true
			}
		}

		if canAccess == false {//跳转404
			w.WriteHeader(403)
			w.Write([]byte("403 Forbidden\n"))
		}
	})

	//beego-orm-adapter
	//a := beegoormadapter.NewAdapter("mysql", beego.AppConfig.String("mysqldsn"), true)
	//e := casbin.NewEnforcer("rbac_model.conf", a)
	//
	//e.LoadPolicy()
	//
	//logs.Info("Policies:", e.GetPolicy())

	//beego.InsertFilter("/admin/*", beego.BeforeRouter, authz.NewAuthorizer(e))

	//authPlugin := auth.NewBasicAuthenticator(SecretAuth, "Authorization Required")
	//beego.InsertFilter("/admin/user", beego.BeforeRouter,authPlugin)
}
