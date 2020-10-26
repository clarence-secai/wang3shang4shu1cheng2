package dao

import (
	"testing"
	//"godeyuandaima4/wang3ye4shu1cheng2/200404/material"
)

// func TestAddBook(t *testing.T){
// 	book := material.Book{Title:"wo",Price:88.8,Sale:55,Stock:5}
// 	AddBook(&book)
// }

// func TestAdd2CartItem(t *testing.T){
// 	uuid :=material.CreateUUID()
// 	book := material.Book{Price:10,Stock:2}
// 	//创建购物项，添加进入数据库
// 	cartitem := material.Cartitem{
// 		CartItemId:uuid,
// 		Book:&book,
// 		Count:1,
// 	}
// 	cartitem.Amount = cartitem.GetAmount()
// 	Add2CartItem(cartitem)
// }

// func TestAdd2Cart(t *testing.T){
// 	uuid :=material.CreateUUID()
// 	book := material.Book{Price:10,Stock:2,Id:110}
// 	//创建购物项，添加进入数据库
// 	cartitem := material.Cartitem{
// 		CartItemId:uuid,
// 		Book:&book,
// 		Count:1,
// 	}
// 	cartitem.Amount = cartitem.GetAmount()
// 	Add2CartItem(cartitem)

// 	//创建购物车，添加进入数据库
// 	var cartitems []*material.Cartitem
// 	cartitems = append(cartitems,&cartitem)

// 	cart := material.Cart{
// 		//Session:session,
// 		CartId:uuid,
// 		UserId:1122,
// 		CartItems:cartitems,
// 	}
// 	cart.TotalCount = cart.GetTotalCount()
// 	cart.TotalAmount = cart.GetTotalAmount()

// 	Add2Cart(&cart)
// }

func TestDeleteAll1(t *testing.T) {
	DeleteAll1(0)
	DeleteAll1(11)
	DeleteAll1(12)

}

// func TestDeleteAll2(t *testing.T){
// 	DeleteAll2(3)

// }
