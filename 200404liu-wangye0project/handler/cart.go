package handler

import (
	"fmt"
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/dao"
	"go2web/wang3ye4shu1cheng2/200404liu-wangye0project/material"
	"html/template"
	"net/http"
	"strconv"
)

//无购物车就新建一个购物车并将书加入到购物车，有购物车就将书加入购物车
func Add2Cart(rw http.ResponseWriter, r *http.Request) {
	session, flag := dao.IsLogin(r)
	if !flag {
		t, _ := template.ParseFiles("./pages/customer/loginpage.html")
		t.Execute(rw, "")
		return
	}
	bookids := r.FormValue("id")
	bookid, _ := strconv.ParseInt(bookids, 10, 64)
	book := dao.GetBookById(bookid) //其实这里也可以从客户端页面上拿到书的信息
	//实际不是book存进数据库中购物项的数据库表，而是book的id存进数据库中购物项的数据库表

	//购物车和购物项通过相同的UUID关联，一个购物车多个购物项，因此先设定UUID，据此
	//来设定后续新增加进来的购物项和购物车
	cart := dao.GetCartByUserId(session.UserId)
	fmt.Println("cart=", cart)
	if cart.CartId == "" { //不要用cart==nil来判断,因为即使无该购物车，查找出来的cart也不会是nil，而是cart= &{<nil>  [] 0 0 0}一些默认值
		//todo:创建购物车和购物项，用UUID关联二者:CartItemId和CartId是相同的，都是UUID
		uuid := material.CreateUUID()
		fmt.Println("打印了cart.go的30行")
		//创建购物项，添加进入数据库
		cartitem := material.Cartitem{
			CartItemId: uuid,
			Book:       book,
			Count:      1,
		}
		cartitem.Amount = cartitem.GetAmount()
		dao.Add2CartItem(cartitem)

		//创建购物车，添加进入数据库
		var cartitems []*material.Cartitem
		cartitems = append(cartitems, &cartitem)

		cart := material.Cart{
			//Session:session,
			CartId:    uuid,
			UserId:    session.UserId,
			CartItems: cartitems,
		}
		cart.TotalCount = cart.GetTotalCount()
		cart.TotalAmount = cart.GetTotalAmount()

		dao.Add2Cart(&cart)
	} else {
		//todo:该用户已经有购物车的情况
		// 首先考虑该书是否已在购物车中，现在是在数量上加一
		flag, count := dao.CheckBook(bookid, session.UserId)
		if flag {
			//说明购物车已有该书一本或若干本，此时再次购买一本相同的
			count++
			cartitem := material.Cartitem{
				CartItemId: cart.CartId, //将已经存在的购物车的UUID赋值给该购物项，从而使得该购物项与该购物车关联
				Book:       book,
				Count:      count,
			}
			cartitem.Amount = cartitem.GetAmount()
			//更新购物项
			dao.UpdateCartItem(&cartitem)

		} else {
			//todo:说明购物车尚无该书，现在是买的第一本
			cartitem := material.Cartitem{
				CartItemId: cart.CartId, //将已经存在的购物车的UUID赋值给该购物项，从而使得该购物项与该购物车关联
				Book:       book,
				Count:      1,
			}
			cartitem.Amount = cartitem.GetAmount()
			//写入到购物项的数据库表
			dao.Add2CartItem(cartitem)
		}
		//更新购物车
		cartitems := dao.GetCartItems(session.UserId)
		dao.UpdateCart(cartitems, session.UserId)
	}
	//更新该图书的销量和库存量
	book.Stock = book.Stock - 1
	book.Sale = book.Sale + 1
	dao.ChangeBookSaleStock(book)
	Index(rw, r)
}

func CartInfo(rw http.ResponseWriter, r *http.Request) {
	session, flag := dao.IsLogin(r)
	if !flag {
		t, _ := template.ParseFiles("./pages/customer/loginpage.html")
		t.Execute(rw, "")
		return
	}
	cartitems := dao.GetCartItems(session.UserId) //这里的购物项里的字段是全的
	cart := dao.GetCartByUserId(session.UserId)   //这里得到的购物车的字段不全
	if cart.CartId != "" {//购物车存在
		realcart := material.Cart{ //其实也不是所有字段都赋值了，但已够客户端页面用
			Session:     session,
			CartId:      cartitems[0].CartItemId,
			CartItems:   cartitems,
			TotalCount:  cart.TotalCount,
			TotalAmount: cart.TotalAmount,
		}
		t, _ := template.ParseFiles("./pages/cart/cart.html")
		t.Execute(rw, realcart)
	} else {//尚未购物而无购物车或已删除购物车的情况
		cart := material.Cart{Session: session}
		t, _ := template.ParseFiles("./pages/cart/cart.html")
		t.Execute(rw, cart)
	}

}

func DeleteCart(rw http.ResponseWriter, r *http.Request) {
	//先要将退回的书的数量还原到数据库和购物页面上
	session, _ := dao.IsLogin(r) //客户端能进入到使用该处理器的页面，肯定是登录了的，否则没机会用这个处理器
	cartitems := dao.GetCartItems(session.UserId)
	for _, v := range cartitems {
		book := material.Book{
			Id: v.Book.Id,
			Title: v.Book.Title,
			Price: v.Book.Price,
			Sale:  v.Book.Sale - v.Count,
			Stock: v.Book.Stock + v.Count}
		dao.ChangeBook(&book)
		//也可以在dao里重新写一个只改book数量的操作数据库函数【update的sql语句中可以只更改数据库
		//表中的某一个字段，数据库表中的其他字段默认的原有值默认保持不变】，上面也无需赋值这么多的book里的字段
	}

	cartid := r.FormValue("cartid")
	dao.DeleteCart(cartid) //删掉数据库中的该用户的购物车和购物项信息

	//todo:调用显示购物车详情的处理器，一个处理器里调用另一个处理器
	CartInfo(rw, r) //在购物车详情页面呈现出空的购物车信息，让用户知道已经删除购物车
	//当用户在购物车详情页面点击返回的index连接后，上面的退回的书的数量就会在
	//客户端购书页面上恢复，因为到index页面所调用的index处理器重新获取数据库里全部书的信息

}
