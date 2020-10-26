package material

import ()

type Page struct {
	UserName  string  //未进入数据库，是由session判断后来赋值的
	PageNo    int64   //当前页的页码
	PageBooks []*Book //每一页中的书
	//pageBooksSize int64 //每页展示图书数量，一般开发者设置好,也可以编写为让用户自定义
	TotalPages int64 //书籍总页数
	IsLogin    bool  //这里和下面的字段均未进入数据库，是在配合模板返回给客户端前给的
	Min        float64
	Max        float64
}

//todo:结构体下的方法，方法名其实也相当于结构体内的字段，
// 注意注意注意：方法名必须开头大写，且结构体实例必须是指针，否则客户端的{{.method}}无法生效

//前一页
func (page *Page) HaveFirst() bool {
	if page.PageNo == 1 {
		return false
	}
	return true
}
func (page *Page) PrePageNo() int64 {
	//	if page.HaveFirst(){
	return page.PageNo - 1
	// }
	//	return 1 //不建议在这里进行判断，在客户端页面进行判断展示最好
}

//后一页
func (page *Page) HaveLast() bool {
	if page.PageNo == page.TotalPages {
		return false
	}
	return true
}
func (page *Page) NextPageNo() int64 {
	return page.PageNo + 1
}
