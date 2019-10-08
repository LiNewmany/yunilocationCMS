package repositories

import (
	"rbacAdmin/models"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
	"database/sql"
	"fmt"
	"strconv"
	"github.com/pkg/errors"
)

/**
 *@author LanguageY++2013
 *2019/2/28 7:29 PM
 **/
func AdminRoles_Pagination(query map[string]string, fields []string, sortby []string, order []string,
	page int64, pageSize int64) (pagination Page, err error){

	offset := (page - 1) * pageSize
	limit := pageSize

	l, err := models.GetAllAdminRoles(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}
	count, err := AdminRoles_GetCount(query)

	pagination = PageUtil(count, page, pageSize, l)

	return
}

func AdminRoles_GetCount(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.AdminRoles))
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

//创建角色
func AdminRoles_CreateRole(name, description string, pIds []int) error {
	o := orm.NewOrm()

	err := o.Begin()
	if err != nil {
		return err
	}

	//创建角色
	var r sql.Result
	r, err = o.Raw("INSERT INTO `admin_roles`(`name`,`display_name`,`description`, `created_at`, `updated_at`) VALUES(?, ?, ?, ?, ?)", name, name, description, time.Now(), time.Now()).Exec()

	var roleId int64
	roleId, err = r.LastInsertId()
	if err != nil {
		o.Rollback()
		return err
	}

	if len(pIds) > 0 {
		//插入权限
		raw_sql := "INSERT INTO `admin_permission_role` VALUES"
		for _, v := range pIds {
			raw_sql += fmt.Sprintf("(%d, %d),", v,roleId)
		}

		raw_sql = strings.TrimRight(raw_sql, ",")

		r, err = o.Raw(raw_sql).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	err = o.Commit()

	if err != nil {
		o.Rollback()
		return err
	}

	return nil
}

//@Description 根据ID删除角色以及对于的权限
func AdminRoles_DelByIds(ids []int) (err error) {
	if len(ids) == 0 {
		return errors.New("ids can not empty")
	}

	idsStr := ""
	for _, v := range ids {
		idsStr += strconv.Itoa(v) + ","
	}

	idsStr = strings.TrimRight(idsStr, ",")

	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return err
	}

	_, err = o.Raw("DELETE FROM admin_roles WHERE id IN(" + idsStr + ")").Exec()
	if err != nil {
		return err
	}

	//删除权限
	_, err = o.Raw("DELETE FROM admin_permission_role WHERE role_id IN(" + idsStr + ")").Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	err = o.Commit()

	if err != nil {
		o.Rollback()
		return err
	}



	return nil
}

//@Description 更新用户角色和权限根据对应的角色ID
func AdminRoles_UpdateById(m *models.AdminRoles, pIds []int)(err error) {
	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return err
	}

	//更新角色
	_, err = o.Raw("UPDATE `admin_roles` SET name=?, display_name=?, description=?, updated_at=? WHERE id=?", m.Name, m.DisplayName, m.Description, m.UpdatedAt, m.Id).Exec()

	//删除角色之前的权限
	_, err = o.Raw("DELETE FROM admin_permission_role WHERE role_id=?", m.Id).Exec()
	if err != nil {
		o.Rollback()
		return err
	}

	//当前的权限值插入
	if len(pIds) > 0 {
		//插入权限
		raw_sql := "INSERT INTO `admin_permission_role` VALUES"
		for _, v := range pIds {
			raw_sql += fmt.Sprintf("(%d, %d),", v, m.Id)
		}

		raw_sql = strings.TrimRight(raw_sql, ",")

		_, err = o.Raw(raw_sql).Exec()
		if err != nil {
			o.Rollback()
			return err
		}
	}

	err = o.Commit()

	if err != nil {
		o.Rollback()
		return err
	}

	return
}