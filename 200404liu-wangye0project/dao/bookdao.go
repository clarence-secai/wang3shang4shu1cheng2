package dao

import (
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/material"
)

func AddBook(book *material.Book) {
	sql := "insert into books(title,price,sale,stock) values (?,?,?,?)"
	material.MyDb.Exec(sql, book.Title, book.Price, book.Sale, book.Stock)
}
func DeleteBook(Id string) {
	sql := "delete from books where id=?"
	material.MyDb.Exec(sql, Id)
}

func ChangeBook(book *material.Book) {
	sql := "update books set title=?,price=?,sale=?,stock=? where id=?"//可以只更改数据库表中某一个字段，其他字段可以会默认不变动
	material.MyDb.Exec(sql, book.Title, book.Price, book.Sale, book.Stock, book.Id)

}
func ChangeBookSaleStock(book *material.Book) {
	sql := "update books set stock=?,sale=? where id=?"
	material.MyDb.Exec(sql, book.Stock, book.Sale, book.Id)

}

func GetBookById(id int64) *material.Book {
	sql := "select * from books where id=?"
	row := material.MyDb.QueryRow(sql, id)
	var book material.Book
	//此时Book结构体里的Num字段被赋值为0，由于当初该字段就未进入
	//数据库，故下一行也没法给该字段赋值，该字段依然是上一行
	//赋给的默认值0
	row.Scan(&book.Id, &book.Title, &book.Price, &book.Sale, &book.Stock)
	return &book
}
func GetBooks() []*material.Book {
	sql := "select * from books "
	rows, _ := material.MyDb.Query(sql)
	var books []*material.Book
	for rows.Next() { //todo:切忌这里是for,不是if
		var book material.Book
		//此时Book结构体里的Num字段被赋值为0，由于当初该字段就未进入
		//数据库，故下一行也没法给该字段赋值，该字段依然是上一行
		//赋给的默认值0
		rows.Scan(&book.Id, &book.Title, &book.Price, &book.Sale, &book.Stock)
		books = append(books, &book)
	}
	return books
}
