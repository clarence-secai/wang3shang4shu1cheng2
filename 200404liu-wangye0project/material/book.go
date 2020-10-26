package material

type Book struct {
	Id    int64
	Title string
	Price float64
	Sale  int64
	Stock int64
	Img   string
	Num   int64 //序号，该字段未进入数据库，辅助在客户端显示数据相对序号
}
