package handler

import (
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/dao"
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/material"
	"html/template"
	"net/http"
	"strconv"
)

//体会：（不跳转页面的情况下）一个页面，必须共用同一个结构体，该结构体包含了该页面所需的所有字段，
//该页面中各连接均指向同一个处理器，处理器返回该结构体和该页面的模板。
//处理器或许不是必须只有一个，比如一个处理器写的内容太多，可以分成多个，但
//结构体必须是同一个。
func Index(rw http.ResponseWriter, r *http.Request) {
	session, flag := dao.IsLogin(r)
	pagenostr := r.FormValue("pageno") //这种也包含了postformvalue的情况
	if pagenostr == "" {
		pagenostr = "1"
	}
	pageno, _ := strconv.ParseInt(pagenostr, 10, 64)

	// mins := r.FormValue("min")
	// if mins==""{
	// 	mins="0"
	// }
	// min,_ := strconv.ParseFloat(mins,64)
	// maxs := r.FormValue("max")
	// if maxs==""{
	// 	maxs="1000"//这里是随便写的一个肯定比数据库中任意一本书单价都高的数字
	// }
	// max,_ := strconv.ParseFloat(maxs,64)
	// //以上更好的方式见如下
	var page *material.Page
	mins := r.FormValue("min")
	maxs := r.FormValue("max")
	if (mins == "" && maxs == "") || (mins == "0" && maxs == "0") { //此处对比看老师的代码，他将Min和Max两个字段设值为string类型的好处
		page = dao.GetPage(pageno)
		//上面这种情况下Max和Min字段未设定，客户端中的{{.Max}}{{.Min}}就显示默认值，没影响
	} else {
		min, _ := strconv.ParseFloat(mins, 64)
		max, _ := strconv.ParseFloat(maxs, 64)

		page = dao.GetPageByPrice(pageno, min, max) //下面两个字段不是从数据库可以取到的，因为原本就没法确定并放进数据库
	}

	if flag == true {
		//session.IsLogin = true
		page.IsLogin = true
		page.UserName = session.UserName
		t, _ := template.ParseFiles("./pages/customer/index.html")
		t.Execute(rw, page)
	} else {
		t, _ := template.ParseFiles("./pages/customer/index.html")
		// session := material.Session{IsLogin:false}
		// //session.IsLogin = false//不能像该行这样，因为else这种情况下session是nil，没法在这里给字段赋值
		page.IsLogin = false
		t.Execute(rw, page)
	}
}
