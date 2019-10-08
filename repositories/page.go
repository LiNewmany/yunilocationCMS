package repositories

/**
 *@author LanguageY++2013
 *2019/2/28 4:33 PM
 **/
//Page分页
type Page struct {
	PageNo		int64
	PageSize	int64
	TotalPage	int64
	TotalCount	int64
	FirstPage	bool
	LastPage	bool
	List 		interface{}
}

func PageUtil(count int64, pageNo int64, pageSize int64, list interface{}) Page {
	tp := count / pageSize
	if count % pageSize > 0 {
		tp = count / pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}

