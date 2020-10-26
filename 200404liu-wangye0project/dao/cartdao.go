package dao

import (
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/material"
)

//在购物车的数据库表中新增一个购物车信息
func Add2Cart(cart *material.Cart) {
	sql := "insert into cart (cartid,totalcount,totalamount,userid) values(?,?,?,?)"
	material.MyDb.Exec(sql, cart.CartId, cart.TotalCount, cart.TotalAmount, cart.UserId)
}

func Add2CartItem(cartitem material.Cartitem) {
	sql := "insert into cartitem (cartItemId,bookid,count,amount) values(?,?,?,?)"
	material.MyDb.Exec(sql, cartitem.CartItemId, cartitem.Book.Id, cartitem.Count, cartitem.Amount)
}

//拿到购物车，实际上此处的购物车只有总数量和总价等四个字段，并不是完整的购物车结构体
func GetCartByUserId(userid int64) *material.Cart {
	sql := "select * from cart where userid=?"
	row := material.MyDb.QueryRow(sql, userid)
	var cart material.Cart
	row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	return &cart
}

//这里获取的购物项结构体是完整的，包括里面的book的完整信息
func GetCartItems(userid int64) []*material.Cartitem {
	sql0 := "select cartid from cart where userid=?"
	row := material.MyDb.QueryRow(sql0, userid)
	var cartid string
	row.Scan(&cartid)

	sql := "select * from cartitem where cartItemId=?"
	rows, _ := material.MyDb.Query(sql, cartid)
	var cartitems []*material.Cartitem
	for rows.Next() {
		var cartitem material.Cartitem
		var bookid int64
		rows.Scan(&cartitem.CartItemId, &bookid, &cartitem.Count, &cartitem.Amount)
		book := GetBookById(bookid)
		cartitem.Book = book

		cartitems = append(cartitems, &cartitem)
	}
	return cartitems
}

func UpdateCart(cartitems []*material.Cartitem, userid int64) {
	var totalcount int64
	var totalamount float64
	for _, v := range cartitems {
		totalcount += v.Count
		totalamount += v.Amount
	}
	sql := "update cart set cartid=?,totalcount=?,totalamount=? where userid=?"
	material.MyDb.Exec(sql, cartitems[0].CartItemId, totalcount, totalamount, userid)
}

func UpdateCartItem(cartitem *material.Cartitem) {
	sql := "update cartitem set count=?,amount=? where bookid=? and cartItemId=?" //此处不能只靠bookid
	material.MyDb.Exec(sql, cartitem.Count, cartitem.Amount, cartitem.Book.Id, cartitem.CartItemId)
}

func DeleteCart(cartid string) {
	sql := "delete from cart where cartid=?"
	material.MyDb.Exec(sql, cartid)
	//也要将相应的购物项删除
	sql = "delete from cartitem where cartItemId=?"
	material.MyDb.Exec(sql, cartid)
}

//查看购物车中是否已经买了一本该书
func CheckBook(bookid int64, userid int64) (bool, int64) {
	sql := "select cartid from cart where userid=?"
	row := material.MyDb.QueryRow(sql, userid)
	var cartid string
	row.Scan(&cartid)
	//其实两个sql语句可以合并成一个
	sql = "select bookid,count from cartitem where cartItemId=?"
	rows, _ := material.MyDb.Query(sql, cartid)
	for rows.Next() {
		var bookidx, count int64
		rows.Scan(&bookidx, &count)
		if bookidx == bookid {
			return true, count
		}
	}
	return false, 0
}
