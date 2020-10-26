package dao

import (
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/material"
	"net/http"
)

func GetUser(username string, password string) *material.User {
	//下一行的and不能用逗号替代,同时写上各个字段或*，不能省略
	sql := "select userid,username,password,email from users where username=? and password=?"
	row := material.MyDb.QueryRow(sql, username, password)
	var user material.User
	row.Scan(&user.UserId, &user.UserName, &user.PassWord, &user.Email)
	return &user
}

func Regist(user *material.User) {
	sql := "insert into users(username,password,email) values (?,?,?)"
	material.MyDb.Exec(sql, user.UserName, user.PassWord, user.Email)

}
func AddSession(session *material.Session) {
	sql := "insert into sessions(sessionid,userid,username) values (?,?,?)"
	material.MyDb.Exec(sql, session.SessionId, session.UserId, session.UserName)

}
func GetSession(uuid string) *material.Session {
	//sql查询语句select和from之间要么是*要么是表头字段，不能没有
	sql := "select * from sessions where sessionid=?"
	row := material.MyDb.QueryRow(sql, uuid)
	var session material.Session
	row.Scan(&session.SessionId, &session.UserId, &session.UserName)
	return &session
}
func DeleteSession(session *material.Session) {
	sql := "delete from sessions where userid=?"
	material.MyDb.Exec(sql, session.UserId)
}
func IsLogin(r *http.Request) (*material.Session, bool) {
	cookie, _ := r.Cookie("usercookie")
	if cookie != nil {
		session := GetSession(cookie.Value)
		if session.UserId > 0 { //对于查数据库后的结果，若非得已，不建议用 ！=nil来判断
			return session, true
		}
	}
	return nil, false
}
