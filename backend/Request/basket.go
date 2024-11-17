package request

type BasketArticleRequest struct {
	UserId    int `json:"userid"`
	ArticleId int `json:"articleid"`
	Quantity  int `json:"quantity"`
}
