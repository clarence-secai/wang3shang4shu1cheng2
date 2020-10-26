package material

type Session struct {
	SessionId string
	//User *User //注意与数据库表的配合，数据库表是不能有自定义类型数据的
	UserId   int64
	UserName string
	IsLogin  bool //该字段未进入数据库表中，默认初始值为false
}

//session是设置与用户匹配的具有独一无二属性信息的字段
