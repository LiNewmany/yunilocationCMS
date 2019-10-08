package controllers

import (
	"rbacAdmin/repositories"
	"rbacAdmin/models"
	"strings"
	"strconv"
	"time"
	"github.com/astaxie/beego/logs"
)

// RbacController operations for Rbac
type RbacController struct {
	BaseController
}

// URLMapping ...
func (c *RbacController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Rbac
// @Param	body		body 	models.Rbac	true		"body for Rbac content"
// @Success 201 {object} models.Rbac
// @Failure 403 body is empty
// @router / [post]
func (c *RbacController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Rbac by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Rbac
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RbacController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Rbac
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Rbac
// @Failure 403
// @router / [get]
func (c *RbacController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Rbac
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Rbac	true		"body for Rbac content"
// @Success 200 {object} models.Rbac
// @Failure 403 :id is not int
// @router /:id [put]
func (c *RbacController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Rbac
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RbacController) Delete() {

}

//@router /permission_list	[get]
func (c *RbacController) PermissionList() {

	query, fields, sortby, order, page, pageSize, err := c.PageParams()
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	//筛选非分类（一级）权限
	query["fid__gt"] = "0"
	pageData, err := repositories.AdminPermissions_Pagination(query, fields, sortby, order, page, pageSize)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	} else {
		//查询分类(一级权限)
		fml, err := models.GetAllAdminPermissions(map[string]string{"fid":"0"}, nil, nil, nil, 0, -1)
		if err == nil {
			c.Data["fml"] = fml
		}

		//重新构建结构
		list := []repositories.SubAdminPermission{}
		Fids := []int{}

		for _, v := range pageData.List.([]interface{}) {
			tmp, ok  := v.(models.AdminPermissions)
			if ok {
				Fids = append(Fids, int(tmp.Fid))
				list = append(list, repositories.SubAdminPermission{
					tmp,
					"",
				})
			}
		}

		//查询分类

		fpl, err := repositories.AdminPermissions_GetByIds(Fids)
		if err == nil {
			for k, v := range list {
				//查询
				if p, ok := fpl[int(v.Fid)]; ok {
					list[k].PName = p.DisplayName
				}
			}
		}

		pageData.List = list

		c.TplName = "rbac/permission_list.html"
		c.Layout = "_layout/iframe_layout.html"
		c.Data["page"] = pageData
	}
}

//@Description	删除
//@Param	ids		string	权限ID 多个ID用逗号分隔
//@router /permission_del	[post]
func(c *RbacController) PermissionDel() {
	idsStr := c.GetString("ids", "")
	if idsStr == "" {
		c.Data["json"] = map[string]interface{}{"code":-1, "msg":"权限ID不能为空"}
		c.ServeJSON()
		return
	}

	ids := []int{}
	id_list := strings.Split(idsStr, ",")
	for _, v := range id_list {
		id, err := strconv.Atoi(v)
		if err == nil {
			ids = append(ids, id)
		}
	}

	if len(ids) == 0 {
		c.Data["json"] = map[string]interface{}{"code":-2, "msg":"没有有效ID"}
		c.ServeJSON()
		return
	}

	err := repositories.AdminPermission_DelByIds(ids)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code":-3, "msg":err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{"code":0, "msg":""}
	c.ServeJSON()
	return
}

//Description	权限编辑
//@router	/permission_edit	[get,post]
func(c *RbacController) PermissionEdit() {
	//获取权限ID
	id, err := c.GetInt32("id")
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	//查询信息
	m, err := models.GetAdminPermissionsById(int(id))
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	if c.Ctx.Request.Method == "GET" {
		//查询分类(一级权限)
		fml, err := models.GetAllAdminPermissions(map[string]string{"fid":"0"}, nil, nil, nil, 0, -1)
		if err == nil {
			c.Data["fml"] = fml
		}

		c.Data["m"] = m
		c.TplName = "rbac/permission_edit.html"
		c.Layout = "_layout/iframe_layout.html"
	}else{
		fid, err := c.GetInt("fid")
		if err != nil {
			c.Data["json"] = err.Error()
			c.ServeJSON()
			return
		}

		m.Fid = uint(fid)
		m.UrlPath = c.GetString("url_path")
		m.DisplayName = c.GetString("display_name")
		m.UpdatedAt = time.Now()

		err = models.UpdateAdminPermissionsById(m)

		if err != nil {
			c.Data["json"] = err.Error()
			c.ServeJSON()
			return
		}

		c.Data["json"] =  map[string]interface{}{"code":0, "msg":""}
		c.ServeJSON()
		return
	}
}

//@Description	权限添加
//@router /permission_add 	[post]
func (c *RbacController) PermissionAdd() {
	fid, err := c.GetInt("fid")
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	m := &models.AdminPermissions{}
	m.Fid = uint(fid)
	m.UrlPath = c.GetString("url_path")
	m.DisplayName = c.GetString("display_name")
	if c.GetString("is_menu") != "" {
		m.IsMenu = 1
	}
	m.UpdatedAt = time.Now()
	m.CreatedAt = time.Now()

	_, err = models.AddAdminPermissions(m)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}else{
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
}

//@router /category_list	[get]
func (c *RbacController) CategoryList() {

	query, fields, sortby, order, page, pageSize, err := c.PageParams()
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	//筛选非分类（一级）权限
	query["fid"] = "0"
	pageData, err := repositories.AdminPermissions_Pagination(query, fields, sortby, order, page, pageSize)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	} else {
		c.TplName = "rbac/category_list.html"
		c.Layout = "_layout/iframe_layout.html"
		c.Data["page"] = pageData
	}
}

//@Description	分类添加
//@router /category_add 	[post]
func (c *RbacController) CategoryAdd() {

	m := &models.AdminPermissions{}
	m.DisplayName = c.GetString("display_name")
	m.IsMenu = 1
	m.UpdatedAt = time.Now()
	m.CreatedAt = time.Now()

	_, err := models.AddAdminPermissions(m)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}else{
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
}

//@Description	删除
//@Param	ids		string	权限ID 多个ID用逗号分隔
//@router /category_del	[post]
func(c *RbacController) CategoryDel() {
	idsStr := c.GetString("ids", "")
	if idsStr == "" {
		c.Data["json"] = map[string]interface{}{"code":-1, "msg":"ID不能为空"}
		c.ServeJSON()
		return
	}

	ids := []int{}
	id_list := strings.Split(idsStr, ",")
	for _, v := range id_list {
		id, err := strconv.Atoi(v)
		if err == nil {
			ids = append(ids, id)
		}
	}

	if len(ids) == 0 {
		c.Data["json"] = map[string]interface{}{"code":-2, "msg":"没有有效ID"}
		c.ServeJSON()
		return
	}

	err := repositories.AdminPermission_DelByIds(ids)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code":-3, "msg":err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{"code":0, "msg":""}
	c.ServeJSON()
	return
}

//Description	权限编辑
//@router	/category_edit	[get,post]
func(c *RbacController) CategoryEdit() {
	//获取权限ID
	id, err := c.GetInt32("id")
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	//查询信息
	m, err := models.GetAdminPermissionsById(int(id))
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	if c.Ctx.Request.Method == "GET" {
		c.Data["m"] = m
		c.TplName = "rbac/category_edit.html"
		c.Layout = "_layout/iframe_layout.html"
	}else{

		m.DisplayName = c.GetString("display_name")
		m.UpdatedAt = time.Now()

		err = models.UpdateAdminPermissionsById(m)

		if err != nil {
			c.Data["json"] = err.Error()
			c.ServeJSON()
			return
		}

		c.Data["json"] =  map[string]interface{}{"code":0, "msg":""}
		c.ServeJSON()
		return
	}
}

//@router /role_list	[get]
func (c *RbacController) RoleList()  {
	query, fields, sortby, order, page, pageSize, err := c.PageParams()

	pageData, err := repositories.AdminRoles_Pagination(query, fields, sortby, order, page, pageSize)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	} else {

		c.TplName = "rbac/role_list.html"
		c.Layout = "_layout/iframe_layout.html"
		c.Data["page"] = pageData
	}
}

//@Description	角色添加
//@router /role_add		[get,post]
func (c *RbacController) RoleAdd() {

	if c.Ctx.Request.Method == "GET" {
		//查询所有权限
		pl, err := models.GetAllAdminPermissions(nil, nil, nil, nil, 0, -1)
		if err == nil {
			//对权限数据按分类进行组织
			cpl := []models.AdminPermissions{}
			for _, v := range pl {
				p := v.(models.AdminPermissions)
				cpl = append(cpl, p)
			}

			cl := repositories.AdminPermission_CategoryTree(cpl)
			c.Data["cl"] = cl
		}

		c.TplName = "rbac/role_add.html"
		c.Layout = "_layout/iframe_layout.html"
	}else{
		ids := []int{}
		params := c.Ctx.Request.Form
		//解析ID
		for k, v := range params {
			if strings.Contains(k, "id[") {
				id, err := strconv.Atoi(v[0])
				if err == nil {
					ids = append(ids, id)
				}
			}
		}


		name := c.GetString("name")
		description := c.GetString("description")

		//创建角色并添加对应的权限
		err := repositories.AdminRoles_CreateRole(name, description, ids)
		if err != nil {
			c.Data["json"] =  map[string]interface{}{"code":-1, "msg":err.Error()}
			c.ServeJSON()
			return
		}

		c.Data["json"] =  map[string]interface{}{"code":0, "msg":""}
		c.ServeJSON()
		return
	}
}

//@Description	角色删除以及对于的权限
//@router	/role_del	[post]
func(c *RbacController) RoleDel() {
	idsStr := c.GetString("ids", "")
	if idsStr == "" {
		c.Data["json"] = map[string]interface{}{"code":-1, "msg":"ID不能为空"}
		c.ServeJSON()
		return
	}

	ids := []int{}
	id_list := strings.Split(idsStr, ",")
	for _, v := range id_list {
		id, err := strconv.Atoi(v)
		if err == nil {
			ids = append(ids, id)
		}
	}

	if len(ids) == 0 {
		c.Data["json"] = map[string]interface{}{"code":-2, "msg":"没有有效ID"}
		c.ServeJSON()
		return
	}

	err := repositories.AdminRoles_DelByIds(ids)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code":-3, "msg":err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{"code":0, "msg":""}
	c.ServeJSON()
	return
}

//@Description	角色编辑
//@router /role_edit		[get,post]
func(c *RbacController) RoleEdit() {
	//获取权限ID
	id, err := c.GetInt32("id")
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	//查询信息
	m, err := models.GetAdminRolesById(int(id))
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	if c.Ctx.Request.Method == "GET" {
		//查询所有权限
		pl, err := models.GetAllAdminPermissions(nil, nil, nil, nil, 0, -1)
		if err == nil {
			//对权限数据按分类进行组织
			cpl := []models.AdminPermissions{}
			for _, v := range pl {
				p := v.(models.AdminPermissions)
				cpl = append(cpl, p)
			}

			cl := repositories.AdminPermission_CategoryTree(cpl)
			c.Data["cl"] = cl
		}

		//查询角色的所有权限
		myPL, err := repositories.RBAC_GetRolePermissions([]int{m.Id})
		if err == nil {
			c.Data["myPL"] = myPL
		}

		c.Data["m"] = m
		c.TplName = "rbac/role_edit.html"
		c.Layout = "_layout/iframe_layout.html"
	}else{

		ids := []int{}
		params := c.Ctx.Request.Form
		//解析ID
		for k, v := range params {
			if strings.Contains(k, "id[") {
				id, err := strconv.Atoi(v[0])
				if err == nil {
					ids = append(ids, id)
				}
			}
		}


		m.Name = c.GetString("name")
		m.Description = c.GetString("description")
		m.DisplayName = m.Name
		m.UpdatedAt = time.Now()

		//创建角色并添加对应的权限m
		err := repositories.AdminRoles_UpdateById(m, ids)
		if err != nil {
			c.Data["json"] =  map[string]interface{}{"code":-1, "msg":err.Error()}
			c.ServeJSON()
			return
		}

		c.Data["json"] =  map[string]interface{}{"code":0, "msg":""}
		c.ServeJSON()
		return
	}
}

//@Description	管理员列表
//@router	/admin_list		[get]
func(c* RbacController) AdminList() {
	query, fields, sortby, order, page, pageSize, err := c.PageParams()

	pageData, err := repositories.AdminUsers_Pagination(query, fields, sortby, order, page, pageSize)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	} else {

		c.TplName = "rbac/admin_list.html"
		c.Layout = "_layout/iframe_layout.html"
		c.Data["page"] = pageData
	}
}

//@Description	管理员添加
//@router	/admin_add		[get,post]
func(c* RbacController) AdminAdd() {
	if c.Ctx.Request.Method == "GET" {
		//查询所有角色
		rl, err := models.GetAllAdminRoles(nil, nil, nil, nil, 0, -1)
		if err == nil {
			c.Data["rl"] = rl
		}

		c.TplName = "rbac/admin_add.html"
		c.Layout = "_layout/iframe_layout.html"
	}else{
		ids := []int{}
		params := c.Ctx.Request.Form
		//解析ID
		for k, v := range params {
			if strings.Contains(k, "id-") {
				id, err := strconv.Atoi(v[0])
				if err == nil {
					ids = append(ids, id)
				}
			}
		}

		name := c.GetString("name")
		mobileNum := c.GetString("mobile_num")
		email := c.GetString("email")
		password := c.GetString("password")
		var isSuper int8
		if c.GetString("is_super") != "" {
			isSuper = 1
		}


		//创建角色并添加对应的权限
		err := repositories.AdminUsers_CreateUser(name, email, mobileNum, password,  isSuper, ids)
		if err != nil {
			c.Data["json"] =  map[string]interface{}{"code":-1, "msg":err.Error()}
			c.ServeJSON()
			return
		}

		c.Data["json"] =  map[string]interface{}{"code":0, "msg":""}
		c.ServeJSON()
		return
	}
}

//@Description	管理员编辑
//@router	/admin_edit	[get,post]
func(c* RbacController) AdminEdit() {
	//查询管理员
	id, err := c.GetInt("id")
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	m, err := models.GetAdminUsersById(id)
	if err != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
		return
	}

	c.Data["m"] = m

	if c.Ctx.Request.Method == "GET" {
		//查询所有角色
		rl, err := models.GetAllAdminRoles(nil, nil, nil, nil, 0, -1)
		if err == nil {
			c.Data["rl"] = rl
		}

		//查询管理员拥有的角色
		roleIds, err := repositories.RBAC_GetUserRoleIds(id)
		logs.Debug(roleIds, err)
		if err == nil {
			c.Data["roleIds"] = roleIds
		}

		c.TplName = "rbac/admin_edit.html"
		c.Layout = "_layout/iframe_layout.html"
	}else{
		ids := []int{}
		params := c.Ctx.Request.Form
		//解析ID
		for k, v := range params {
			if strings.Contains(k, "id-") {
				id, err := strconv.Atoi(v[0])
				if err == nil {
					ids = append(ids, id)
				}
			}
		}

		name := c.GetString("name")
		mobileNum := c.GetString("mobile_num")
		email := c.GetString("email")
		password := c.GetString("password")
		var isSuper int8
		if c.GetString("is_super") != "" {
			isSuper = 1
		}


		//创建角色并添加对应的权限
		err := repositories.AdminUsers_UpdateUser(id, name, email, mobileNum, password,  isSuper, ids)
		if err != nil {
			c.Data["json"] =  map[string]interface{}{"code":-1, "msg":err.Error()}
			c.ServeJSON()
			return
		}

		c.Data["json"] =  map[string]interface{}{"code":0, "msg":""}
		c.ServeJSON()
		return
	}
}

//@Description	管理员(批量)删除
//@router	/admin_del		[post]
func(c* RbacController) AdminDel() {
	idsStr := c.GetString("ids", "")
	if idsStr == "" {
		c.Data["json"] = map[string]interface{}{"code":-1, "msg":"ID不能为空"}
		c.ServeJSON()
		return
	}

	ids := []int{}
	id_list := strings.Split(idsStr, ",")
	for _, v := range id_list {
		id, err := strconv.Atoi(v)
		if err == nil {
			ids = append(ids, id)
		}
	}

	if len(ids) == 0 {
		c.Data["json"] = map[string]interface{}{"code":-2, "msg":"没有有效ID"}
		c.ServeJSON()
		return
	}

	err := repositories.AdminUsers_DelByIds(ids)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code":-3, "msg":err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{"code":0, "msg":""}
	c.ServeJSON()
	return
}



