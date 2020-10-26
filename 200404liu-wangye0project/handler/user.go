package handler

import (
	"fmt"
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/dao"
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/material"
	"html/template"
	"net/http"
)

func Regist(rw http.ResponseWriter, r *http.Request) {
	username := r.FormValue("yonghuming") //不用PostFormValue也可以

	fmt.Println("username=", username)
	password := r.PostFormValue("mima")
	email := r.PostFormValue("youxiang")
	user := dao.GetUser(username, password)

	fmt.Println("user=", user) //发现这里会打印&{0}故下一行不能用 ！= nil判断
	if user.Email == email {   //表明该用户已存在，无需再注册
		fmt.Println("用户已存在，已注册")
		t, _ := template.ParseFiles("./pages/customer/registfail.html")
		t.Execute(rw, user)
	} else { //需向数据库注册该用户
		user := material.User{UserName: username, PassWord: password, Email: email}
		dao.Regist(&user)
		t, _ := template.ParseFiles("./pages/customer/registsuccess.html")
		t.Execute(rw, user)
	}
}

func Login(rw http.ResponseWriter, r *http.Request) {
	//为了避免登录界面的重复刷新以致重复执行本处理器
	//导致在服务器产生多余session，可以先判断IsLogin
	session, flag := dao.IsLogin(r)
	if flag == true {
		session.IsLogin = true
		t, _ := template.ParseFiles("./pages/customer/index.html")
		t.Execute(rw, session)//todo:这里还是有问题，如果登录时重复刷新，则返回的页面会不完整
		return
	}
	username := r.FormValue("yonghuming")
	password := r.FormValue("mima")
	user := dao.GetUser(username, password) //会发现不区分大小写，比如登录信息写clarence，也会返回数据库里的Clarence
	//与17行一样，下面一行同样不能以！=nil判断
	if user.UserId > 0 { //表明该用户已存在，登录成功
		// t,_ := template.ParseFiles("./pages/customer/loginsuccess.html")
		// t.Execute(rw,user)
//todo:注意：上面两行代码只能在62行http.SetCookie后面否则cookie的设置无法完成！
	//todo:下面开始设置cookie和session,二者以拥有相同的UUID来实现一一对应
		uuid := material.CreateUUID()
		session := material.Session{
			SessionId: uuid,
			//User : user,//上面返回的user已经是指针类型
			UserId:   user.UserId,
			UserName: user.UserName,
			//IsLogin:true,//该字段并未存进数据库表
		}
		dao.AddSession(&session)
		mycookie := http.Cookie{
			Name:  "usercookie",
			Value: uuid,
		}
		http.SetCookie(rw, &mycookie)
		t, _ := template.ParseFiles("./pages/customer/loginsuccess.html")
		t.Execute(rw, user)
	} else { //该用户还没注册或用户名密码不对
		t, _ := template.ParseFiles("./pages/customer/loginfail.html")
		t.Execute(rw, "")
	}
}

func Logout(rw http.ResponseWriter, r *http.Request) {
	session, _ := dao.IsLogin(r)
	dao.DeleteSession(session) //只删除数据库的session还不够，还需要让浏览器的cookie失效
	nulcookie := http.Cookie{MaxAge: -1}
	http.SetCookie(rw, &nulcookie)
	Index(rw, r)
}
