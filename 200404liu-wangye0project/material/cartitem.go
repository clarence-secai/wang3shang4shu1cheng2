package material

import ()

type Cartitem struct {
	CartItemId string  //购物项id与购物车id相同,都是同一个UUID
	Book       *Book   //也可为BOOK里的各字段，但这样继承更好，实际存入数据库的是book的id
	Count      int64   //购买某本书的数量
	Amount     float64 //购买该书的总价,该字段需借助Count字段和下面的方法求得
}

func (cartitem *Cartitem) GetAmount() float64 {
	return cartitem.Book.Price * float64(cartitem.Count)
}
