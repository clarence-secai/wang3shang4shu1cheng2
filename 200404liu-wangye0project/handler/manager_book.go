package handler

import (
	"fmt"
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/dao"
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/material"
	"html/template"
	"net/http"
	"strconv"
)

//显示所有书
func BookPage(rw http.ResponseWriter, r *http.Request) {
	books := dao.GetBooks()
	for k, v := range books {
		//todo:给显示到前端的全部书籍的进行编号
		v.Num = int64(k) + 1 //这里不能再 := 因Num这个字段本身一直存在，且在上一行中拿到的book中该字段已经有默认赋值0
	}
	t, _ := template.ParseFiles("./pages/manager/bookpage.html")
	t.Execute(rw, books)
}
//把要添加或修改一本书的书的信息获取到，并转到一个专门用来填入书的新信息的页面来向数据库提交
func BookBoxPage(rw http.ResponseWriter, r *http.Request) {
	sid := r.FormValue("id") //注意这里不能用PostFormValue
	if sid == "" { //即是前端点击的是添加一本书的按钮
		book := material.Book{}
		t, _ := template.ParseFiles("./pages/manager/bookbox.html")
		t.Execute(rw, book) //这里会导致book.Id默认给了0
	} else {//即前端点击的是修改一本书信息的按钮
		Id, _ := strconv.ParseInt(sid, 10, 64)
		stitle := r.PostFormValue("shuming")
		sprice := r.PostFormValue("jiage")
		price, _ := strconv.ParseFloat(sprice, 64)
		ssale := r.PostFormValue("xiaoliang")
		sale, _ := strconv.ParseInt(ssale, 10, 64)
		sstock := r.PostFormValue("kucun")
		stock, _ := strconv.ParseInt(sstock, 10, 64)
		book := material.Book{Id: Id, Title: stitle, Price: price, Sale: sale, Stock: stock}
		t, _ := template.ParseFiles("./pages/manager/bookbox.html")
		t.Execute(rw, book)
	}
}
//把填入书的新信息的页面向数据库提交
func AddOrChangeBook(rw http.ResponseWriter, r *http.Request) {
	sid := r.FormValue("id")
	id, _ := strconv.ParseInt(sid, 10, 64)

	stitle := r.PostFormValue("shuming")

	sprice := r.PostFormValue("jiage")
	price, _ := strconv.ParseFloat(sprice, 64)

	ssale := r.PostFormValue("xiaoliang")
	sale, _ := strconv.ParseInt(ssale, 10, 64)

	sstock := r.PostFormValue("kucun")
	stock, _ := strconv.ParseInt(sstock, 10, 64)
	//下面的book1写上Id:id，即赋值默认的0值，不影响数据库id自增
	book1 := material.Book{Title: stitle, Price: price, Sale: sale, Stock: stock}
	//此时Book结构体里的字段Num是默认赋的0，该字段未进入数据库
	book2 := material.Book{Id: id, Title: stitle, Price: price, Sale: sale, Stock: stock}
	if id == 0 { //这里不应该是sid=="",而是sid==0
		fmt.Println("打印了我") //用于找错，看该处及以后是否运行了
		dao.AddBook(&book1)
	} else {
		//book.Id = id
		fmt.Println("打印了wo")
		dao.ChangeBook(&book2)
	}
	//todo:调用另一个处理器  一个处理器中调用另一个处理器
	BookPage(rw, r)
}

//从显示的所有书中删掉一条
func DeleteBook(rw http.ResponseWriter, r *http.Request) {
	sid := r.FormValue("id")
	//Id,_:= strconv.ParseInt(sid,10,64)
	dao.DeleteBook(sid)
	//todo:调用另一个处理器  一个处理器里调用另一个处理器
	BookPage(rw, r) //这样才能在客户端体现删掉了一本书
}
