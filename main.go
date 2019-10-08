package main

import (
	_ "rbacAdmin/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	_ "rbacAdmin/filters"
	"rbacAdmin/views"
)


func init() {


	//ACL
	//beego.InsertFilter("~/login", beego.BeforeRouter, authz.NewAuthorizer(casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")), false)



	//e.AddPermissionForUser("alice", "/user")
	//e.AddGroupingPolicy("guest", "/user/*", "*")

	//为角色guest添加策略
	//e.AddPolicy("guest", "/user/*", "*")
	//e.AddPolicy("guest", "/teacher/*", "*")
	//e.AddPolicy("guest", "/course/*", "*")
	//e.AddPolicy("guest", "/video/*", "*")

	//为用户alice赋予guest角色
	//e.AddRoleForUser("alice","guest")

	//e.SavePolicy()


	//注册mysql

	orm.RegisterDriver("mysql", orm.DRMySQL)
	// dsn := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@tcp(" +beego.AppConfig.String("mysqlurls") +")/" + beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("mysqldsn"))

	orm.Debug = true

	//添加模板函数
	beego.AddFuncMap("CheckPermission", views.CheckPermission)
	beego.AddFuncMap("IsContains", views.IsContains)
}

func main() {
	beego.Run()
}

