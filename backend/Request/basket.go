package request

import entity "backend/Entity"

type FullBasket struct {
	UserId     int            `json:"userid"`
	Articles   []BasketDetail `json:"articles"`
	TotalPrice float32        `json:"totalprice"`
}

type BasketDetail struct {
	ArticleDetail entity.Article `json:"article"`
	Quantity      int            `json:"quantity"`
	LinePrice     float32        `json:"lineprice"`
}

type BasketArticleRequest struct {
	UserId    int `json:"userid"`
	ArticleId int `json:"articleid"`
	Quantity  int `json:"quantity"`
}
