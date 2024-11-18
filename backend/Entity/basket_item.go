package entity

type BasketItem struct {
	UserId    int `json:"userid" gorm:"primaryKey"`
	ArticleId int `json:"articleid" gorm:"primaryKey"`
	Quantity  int `json:"quantity"`
}
