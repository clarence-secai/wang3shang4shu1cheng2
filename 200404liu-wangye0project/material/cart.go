package material

type Cart struct {
	Session     *Session  //未保存进数据库，因为同一个客户不同次登陆，session都不一样
	CartId      string //由UUID产生，从而开发者可以将其同样给到对应的CartItemId，如果由数据库去自增就没法这么做
	CartItems   []*Cartitem  //未进入数据库保存，也没法保存
	TotalCount  int64 //两个字段需借助CartItems字段和下面两个方法获得
	TotalAmount float64
	UserId      int64 //与用户一对一关联
}

func (cart *Cart) GetTotalCount() int64 {
	var sumcount int64
	for _, v := range cart.CartItems {
		sumcount += v.Count
	}
	return sumcount
}
func (cart *Cart) GetTotalAmount() float64 {
	var sumamount float64
	for _, v := range cart.CartItems {
		//sumamount += v.Amount
		//todo:错误，这里极其容易出错，除非调用该方法的cart里的cartitem里的Amount字段已赋值，无需再计算去的Amount
		sumamount += v.GetAmount()
	}
	return sumamount
}
