package request

type ArticleRequest struct {
	ArticleId   int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImgUrl      string  `json:"imgurl"`
}

type ArticleLine struct {
	ArticleId int     `json:"articleid"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
}
