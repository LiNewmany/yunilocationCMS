package repositories

import (
	"rbacAdmin/models"
	"github.com/astaxie/beego/orm"
	"strings"
	"github.com/pkg/errors"
	"strconv"
)

/**
 *@author LanguageY++2013
 *2019/2/28 7:29 PM
 **/
func AdminPermissions_Pagination(query map[string]string, fields []string, sortby []string, order []string,
page int64, pageSize int64) (pagination Page, err error){

	offset := (page - 1) * pageSize
	limit := pageSize

	l, err := models.GetAllAdminPermissions(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}
	count, err := AdminPermissions_GetCount(query)

	pagination = PageUtil(count, page, pageSize, l)

	return
}

func AdminPermissions_GetCount(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.AdminPermissions))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}


	count, err =  qs.Count()
	return
}


type SubAdminPermission struct {
	models.AdminPermissions
	PName		string	//所属分类名称
}


func AdminPermissions_GetByIds(ids []int)(map[int]models.AdminPermissions, error){
	if len(ids) == 0{
		return nil, errors.New("ids is empty")
	}

	ids_tmp := []interface{}{}
	for _, id := range ids {
		ids_tmp = append(ids_tmp, id)
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.AdminPermissions))

	l := []models.AdminPermissions{}
	_, err := qs.Filter("id__in", ids_tmp...).All(&l)
	if err != nil {
		return nil, err
	}

	m := map[int]models.AdminPermissions{}
	for _, v := range l {
		m[v.Id] = v
	}

	return m, nil
}


func AdminPermission_DelByIds(ids []int)(err error) {
	if len(ids) == 0 {
		return errors.New("ids can not empty")
	}

	idsStr := ""
	for _, v := range ids {
		idsStr += strconv.Itoa(v) + ","
	}

	idsStr = strings.TrimRight(idsStr, ",")

	o := orm.NewOrm()
	_, err = o.Raw("DELETE FROM admin_permissions WHERE id IN(" + idsStr + ")").Exec()

	return
}

type CategoryNode struct {
	Self models.AdminPermissions
	SubNodes	[]models.AdminPermissions
}

//构建2维权限列表
func AdminPermission_CategoryTree(permissions []models.AdminPermissions)(categories map[int]*CategoryNode) {
	categories = map[int]*CategoryNode{}

	//首先找到分类
	for _, v := range permissions  {
		if v.Fid == 0 {
			m := &CategoryNode{
				Self:v,
				SubNodes:[]models.AdminPermissions{},
			}

			categories[v.Id] = m
		}
	}

	//二级菜单填充
	for _, v := range permissions {
		if m, ok := categories[int(v.Fid)]; ok {
			m.SubNodes = append(m.SubNodes, v)
		}
	}

	return
}