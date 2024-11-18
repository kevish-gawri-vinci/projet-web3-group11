package entity

type ArticleLine struct {
	ArticleId int     `json:"articleid"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
}
