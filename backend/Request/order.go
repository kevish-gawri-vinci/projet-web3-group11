package request

import entity "backend/Entity"

type FullOrder struct {
	OrderId    int           `json:"orderid"`
	Articles   []OrderDetail `json:"articles"`
	TotalPrice float32       `json:"totalprice"`
}

type OrderDetail struct {
	ArticleLine   ArticleLine    `json:"articleline"`
	ArticleDetail entity.Article `json:"articledetail"`
}

//For /order/getall ==> Returning a list of orders
type OrderListLine struct {
	OrderId       int     `json:"orderid"`
	TotalPrice    float32 `json:"totalprice"`
	TotalQuantity int     `json:"totalquantity"`
}
