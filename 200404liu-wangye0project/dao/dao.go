package dao

import (
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/material"
)

func DeleteAll1(userid int64) {
	sql := "delete from cart where userid=?"
	material.MyDb.Exec(sql, userid)
}
func DeleteAll2(bookid int64) {
	sql := "delete from cartitem where bookid=?"
	material.MyDb.Exec(sql, bookid)
}
