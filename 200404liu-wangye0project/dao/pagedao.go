package dao

import (
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/material"
	"math"
)

//type Page struct{
// 	UserName string
// 	PageNo int64 //当前页的页码
// 	PageBooks []*Book//每一页中的书
// 	//pageBooksSize int64 //每页展示图书数量，一般开发者设置好,也可以编写为让用户自定义
// 	TotalPages int64//书籍总页数
// 	IsLogin bool
// 	Img string
// }
//拿到有价格限定的图书分页
func GetPage(pageno int64) *material.Page {
	var pageBooksSize int64 = 4         //每页展示4本图书
	sql := "select count(*) from books" //像这样的就无需where
	var count int64
	row := material.MyDb.QueryRow(sql) //也可以没后面的参数，使用的也是QueryRow
	row.Scan(&count)                   //获得数据库图书记录总条数
	totalpages := math.Ceil(float64(count) / float64(pageBooksSize))

	sql2 := "select * from books limit ?,?" //limit的是记录条数
	rows, _ := material.MyDb.Query(sql2, pageBooksSize*(pageno-1), pageBooksSize)
	var books []*material.Book
	for rows.Next() {
		var book material.Book
		rows.Scan(&book.Id, &book.Title, &book.Price, &book.Sale, &book.Stock)
		book.Img = "/picture/book.jpg"
		books = append(books, &book)
	}
	//下面只给Page结构体的三个字段赋值了
	page := material.Page{TotalPages: int64(totalpages), PageNo: pageno, PageBooks: books}
	return &page
}

func GetPageByPrice(pageno int64, min float64, max float64) *material.Page {
	var pageBooksSize int64 = 4 //每页展示4本图书
	sql := "select count(*) from books where price between ? and ?"
	var count int64
	row := material.MyDb.QueryRow(sql, min, max)
	row.Scan(&count) //获得数据库图书记录总条数
	totalpages := math.Ceil(float64(count) / float64(pageBooksSize))

	sql2 := "select * from books where price between ? and ? limit ?,?" //limit的是记录条数
	rows, _ := material.MyDb.Query(sql2, min, max, pageBooksSize*(pageno-1), pageBooksSize)
	var books []*material.Book
	for rows.Next() {
		var book material.Book
		rows.Scan(&book.Id, &book.Title, &book.Price, &book.Sale, &book.Stock)
		book.Img = "/picture/book.jpg"  //todo:将图片链接字符串赋值给字段便于使用
		books = append(books, &book)
	}
	//下面只给Page结构体的五个字段赋了值
	page := material.Page{
		TotalPages: int64(totalpages),
		PageNo: pageno,
		PageBooks: books,
		Min: min,
		Max: max}
	return &page
}
