package request

type AddOneArticleRequest struct {
	UserId    int `json:"userid"`
	ArticleId int `json:"articleid"`
	Quantity  int `json:"quantity"`
}
