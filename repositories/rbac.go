package repositories

import (
	"rbacAdmin/models"
	"github.com/astaxie/beego/orm"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"github.com/astaxie/beego/logs"
)

/**
 *@author LanguageY++2013
 *2019/2/28 10:13 AM
 **/
func RBAC_GetUserPermissions(userId int)(permissions []models.AdminPermissions, err error) {
	//判断是否是超级用户
	var v *models.AdminUsers
	v, err = models.GetAdminUsersById(userId)
	if err != nil {
		return
	}

	if v.IsSuper == 1 {//超级管理员，获取所有权限
		var ml []interface{}
		ml, err = models.GetAllAdminPermissions(nil,nil, nil, nil, 0, -1 )
		if err != nil {
			return
		}

		for _, li := range ml   {
			permission :=li.(models.AdminPermissions)
			permissions = append(permissions, permission)
		}

		return
	}

	//非超级管理员，根据用户角色获取权限
	roleIds, err := RBAC_GetUserRoleIds(userId)
	if err == nil && len(roleIds) > 0 {//拥有角色
		//查询角色拥有的权限
		permissions, err = RBAC_GetRolePermissions(roleIds)
	}

	return
}

/**
 *@description 获取用户角色
 *@params	userId	int 	用户ID
 */
func RBAC_GetUserRoleIds(userId int)(roleIds []int, err error) {
	o := orm.NewOrm()

	var list orm.ParamsList
	_, err = o.Raw("SELECT role_id FROM `admin_role_user` WHERE admin_user_id = ?", userId).ValuesFlat(&list)
	if err != nil {
		return
	}

	logs.Debug("List:", list)

	for _, v := range list {
		s, ok := v.(string)
		if ok {
			id, _ := strconv.Atoi(s)
			roleIds = append(roleIds, id)
		}
	}

	return
}

/**
 *@description 获取角色拥有的权限
 */
func RBAC_GetRolePermissions(roleIds []int)(permission []models.AdminPermissions, err error) {
	if len(roleIds) <= 0 {
		err = errors.New("role ids length must larger than 0")
		return
	}

	inStr := ""
	for _, id := range roleIds {
		inStr += strconv.Itoa(id) + ","
	}

	inStr = strings.TrimRight(inStr, ",")

	o := orm.NewOrm()
	_, err = o.Raw("SELECT p.* FROM admin_permissions AS p LEFT JOIN admin_permission_role pr ON pr.`permission_id`=p.`id` WHERE pr.`role_id` IN(?)", inStr).QueryRows(&permission)

	return
}

type Menus struct {
	Name 	string
	Icon 	string
	Url 	string
	SubMenus []Menus
}
/**
 *@description 对权限数据进行处理 筛选出菜单 并形成二层菜单结构
 */
 func RBAC_Menus(permissions []models.AdminPermissions)(menus map[int]*Menus) {
	menus = map[int]*Menus{}

	//首先找到一级菜单
	 for _, v := range permissions  {
	 	if v.IsMenu == 0 {
	 		continue
		}

	 	if v.Fid == 0 {
			m := &Menus{
				Name:v.DisplayName,
				Url:v.UrlPath,
				Icon:v.Icon,
				SubMenus:[]Menus{},
			}

			menus[v.Id] = m
		}
	 }

	 //二级菜单填充
	 for _, v := range permissions {
	 	if v.IsMenu == 0 || v.Fid == 0 {
	 		continue
		}

	 	if m, ok := menus[int(v.Fid)]; ok {
			m.SubMenus = append(m.SubMenus, Menus{Name:v.DisplayName, Url:v.UrlPath, Icon:v.Icon})
		}
	 }

	 return
 }

