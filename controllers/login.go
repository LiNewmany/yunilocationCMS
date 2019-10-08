package controllers

import (
	"github.com/astaxie/beego"
	"rbacAdmin/models"
	"golang.org/x/crypto/bcrypt"
	"rbacAdmin/repositories"
	"github.com/astaxie/beego/logs"
)

// LoginController operations for Login
type LoginController struct {
	beego.Controller
}

// URLMapping ...
func (c *LoginController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Login
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 201 {object} models.Login
// @Failure 403 body is empty
// @router / [post]
func (c *LoginController) Post() {

}


// GetAll ...
// @Title GetAll
// @Description get Login
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Login
// @Failure 403
// @router / [get]
func (c *LoginController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Login
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Login	true		"body for Login content"
// @Success 200 {object} models.Login
// @Failure 403 :id is not int
// @router /:id [put]
func (c *LoginController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Login
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *LoginController) Delete() {

}

// @router /login	[get,post]
func (c *LoginController) Login() {
	if c.Ctx.Request.Method == "GET" {
		c.TplName = "login/login.html"
	}else{
		var email, password string
		email = c.GetString("email")
		password = c.GetString("password")

		//获取用户名，密码
		ml, err := models.GetAllAdminUsers(map[string]string{"email":email}, nil, nil, nil, 0, 1)
		if err != nil || len(ml) <= 0{
			c.Data["json"] = map[string]interface{}{"code":-1, "msg":"用户信息查询失败"}
			c.ServeJSON()
			return
		}

		user := ml[0].(models.AdminUsers)

		if  bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
			c.Data["json"] = map[string]interface{}{"code":-2, "msg":"密码不正确"}

			//bcryptSecret, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
			//if err == nil {
			//	user.Password = string(bcryptSecret)
			//
			//	models.UpdateAdminUsersById(&user)
			//}

			c.ServeJSON()
			return
		}

		//c.SetSession("role", c.GetString("role"))
		//c.SetSession("username", c.GetString("username"))
		//c.SetSession("password", c.GetString("password"))

		c.SetSession("user", user)

		//跳转首页
		//c.Ctx.Redirect(302, "/admin/user/")

		//权限数据测试
		pl, err := repositories.RBAC_GetUserPermissions(user.Id)
		if err == nil {
			logs.Info(pl)
			c.SetSession("permissions", pl)
			//左侧菜单
			menus := repositories.RBAC_Menus(pl)
			logs.Info(menus)
			c.SetSession("menus", menus)
		}

		c.Data["json"] = map[string]interface{}{"code":0, "msg":""}
		c.ServeJSON()
	}
}

// @router login_out	[get]
func (c *LoginController) LoginOut() {

	c.DelSession("role")
	c.DelSession("username")
	c.DelSession("password")



	c.Ctx.Redirect(302, "/admin/login/login")
}
