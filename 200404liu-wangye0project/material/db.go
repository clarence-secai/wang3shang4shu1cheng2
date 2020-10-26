package material

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var MyDb *sql.DB
var err error //error是一个数据类型，就像string、float32等一样
func init() { //下面一行切忌不能是 :=
	MyDb, err = sql.Open("mysql", "root:413188ok@tcp(localhost:3306)/test")
	fmt.Println(MyDb.Ping())
	if err != nil {
		fmt.Println("打开数据库出错", err)
		return
	}
}
