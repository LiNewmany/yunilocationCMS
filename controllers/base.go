package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"github.com/pkg/errors"
	"github.com/astaxie/beego/logs"
)

/**
 *@author LanguageY++2013
 *2019/2/28 8:58 PM
 **/
const (
	DEFAULT_PAGE = 1
	DEFAULT_PAGESIZE = 10
)

type BaseController struct {
	beego.Controller
}

func(c *BaseController) PageParams() (query map[string]string, fields []string, sortby []string, order []string,
	page int64, pageSize int64, err error){

	//获取所有二级权限
	query = make(map[string]string)


	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				err = errors.New("Error: invalid query key/value pair")
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	params := c.Ctx.Request.Form
	for k, v := range params {
		if len(v) == 1 && v[0] != "" && strings.Contains(k, "page") == false {
			query[k] = v[0]
		}
	}

	logs.Info("query", query)

	// limit: 10 (default is 10)
	if v, err := c.GetInt64("page"); err == nil {
		page = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("pageSize"); err == nil {
		pageSize = v
	}

	if pageSize == 0 {
		pageSize = DEFAULT_PAGESIZE
	}

	if page == 0 {
		page = DEFAULT_PAGE
	}

	return

}