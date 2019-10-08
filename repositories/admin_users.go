package repositories

import (
	"github.com/astaxie/beego/orm"
	"strings"
	"rbacAdmin/models"
	"golang.org/x/crypto/bcrypt"
	"time"
	"database/sql"
	"fmt"
	"strconv"
	"github.com/pkg/errors"
)

/**
 *@author LanguageY++2013
 *2019/2/28 4:38 PM
 **/
func GetAllAdminUsersCount(query map[string]string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.AdminUsers))
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

func AdminUsers_Pagination(query map[string]string, fields []string, sortby []string, order []string,
	page int64, pageSize int64) (pagination Page, err error){

	offset := (page - 1) * pageSize
	limit := pageSize

	l, err := models.GetAllAdminUsers(query, fields, sortby, order, offset, limit)
	if err != nil {
		return
	}
	count, err := GetAllAdminUsersCount(query)

	pagination = PageUtil(count, page, pageSize, l)

	return
}

//创建角色并且赋予对应的角色
func AdminUsers_CreateUser(name, email, mobile_num,  password string, isSuper int8, rIds []int) (err error) {

	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	hashPasswordT := []byte{}
	hashPasswordT, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return
	}

	hashPassword := string(hashPasswordT)

	var r sql.Result
	r, err = o.Raw("INSERT INTO `admin_users`(`name`,`email`,`mobile_num`,`password`,`is_super`,`created_at`,`updated_at`) VALUES(?, ?, ?, ?, ?, ?, ?)", name, email, mobile_num, hashPassword, isSuper, time.Now(), time.Now()).Exec()
	if err != nil {
		return
	}

	userId, err := r.LastInsertId()
	if err != nil{
		return
	}

	//赋予角色
	if len(rIds) > 0 {

		ur_sql := "INSERT INTO `admin_role_user`(`admin_user_id`,`role_id`) VALUES"
		for _, role_id := range rIds {
			ur_sql += fmt.Sprintf("('%d','%d'),", userId, role_id)
		}

		ur_sql = strings.TrimRight(ur_sql, ",")

		//执行插入
		r, err = o.Raw(ur_sql).Exec()
		if err != nil {
			o.Rollback()
			return
		}
	}

	err = o.Commit()

	if err != nil {
		o.Rollback()
		return
	}

	return
}

//创建角色并且赋予对应的角色
func AdminUsers_UpdateUser(admin_user_id int, name, email, mobile_num,  password string, isSuper int8, rIds []int) (err error) {

	o := orm.NewOrm()

	err = o.Begin()
	if err != nil {
		return
	}

	if password != "" {
		hashPasswordT := []byte{}
		hashPasswordT, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			return
		}

		hashPassword := string(hashPasswordT)

		_, err = o.Raw("UPDATE `admin_users` SET `name`=?,`email`=?,`mobile_num`=?,`password`=?,`is_super`=?,`updated_at`=? WHERE id=?", name, email, mobile_num, hashPassword, isSuper, time.Now(), admin_user_id).Exec()
		if err != nil {
			return
		}
	}else{
		_, err = o.Raw("UPDATE `admin_users` SET `name`=?,`email`=?,`mobile_num`=?,`is_super`=?,`updated_at`=? WHERE id=?", name, email, mobile_num,  isSuper, time.Now(), admin_user_id).Exec()
		if err != nil {
			return
		}
	}


	//删除旧有角色
	_, err = o.Raw("DELETE FROM admin_role_user WHERE admin_user_id=?",  admin_user_id).Exec()
	if err != nil {
		o.Rollback()
		return
	}


	//赋予角色
	if len(rIds) > 0 {

		ur_sql := "INSERT INTO `admin_role_user`(`admin_user_id`,`role_id`) VALUES"
		for _, role_id := range rIds {
			ur_sql += fmt.Sprintf("('%d','%d'),", admin_user_id, role_id)
		}

		ur_sql = strings.TrimRight(ur_sql, ",")

		//执行插入
		_, err = o.Raw(ur_sql).Exec()
		if err != nil {
			o.Rollback()
			return
		}
	}

	err = o.Commit()

	if err != nil {
		o.Rollback()
		return
	}

	return
}

//删除用户及对应的角色
func AdminUsers_DelByIds(ids []int) (err error) {
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

	_, err = o.Raw("DELETE FROM admin_users WHERE id IN(" + idsStr + ")").Exec()
	if err != nil {
		return err
	}

	//删除权限
	_, err = o.Raw("DELETE FROM admin_role_user WHERE admin_user_id IN(" + idsStr + ")").Exec()
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