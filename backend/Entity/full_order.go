package entity

type FullOrder struct {
	OrderId    int           `json:"orderid"`
	Articles   []OrderDetail `json:"articles"`
	TotalPrice float32       `json:"totalprice"`
}

type OrderDetail struct {
	ArticleLine   ArticleLine `json:"articleline"`
	ArticleDetail Article     `json:"articledetail"`
}
