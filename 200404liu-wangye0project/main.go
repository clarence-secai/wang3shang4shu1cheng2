package main

import (
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/index", handler.Index)
	http.Handle("/picture/", http.StripPrefix("/picture/", http.FileServer(http.Dir("./picture"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("./pages"))))
	//上面一行对应index.html里的href链接，这种方式一般不太常用，因为遇到页面里包含{{.}}情况的，这种没有template的直接进页面会呈现{{.}}
	//下面是管理员管理的所有书籍的增删改
	http.HandleFunc("/bookpage", handler.BookPage)
	http.HandleFunc("/deletebook", handler.DeleteBook)
	http.HandleFunc("/bookboxpage", handler.BookBoxPage)
	http.HandleFunc("/addorchangebook", handler.AddOrChangeBook) //拼写出错客户端会出现404，原因是拼写出错没能调用处理器
	//注册和登录操作
	http.HandleFunc("/regist", handler.Regist)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/logout", handler.Logout)
	//添加购物车操作
	http.HandleFunc("/add2cart", handler.Add2Cart)
	http.HandleFunc("/cartinfo", handler.CartInfo)
	http.HandleFunc("/deletecart", handler.DeleteCart)

	http.ListenAndServe(":8888", nil) //这里必须有冒号
}
