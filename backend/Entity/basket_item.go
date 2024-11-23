package entity

type BasketItem struct {
	UserId    int `gorm:"primaryKey"`
	ArticleId int `gorm:"primaryKey"`
	Quantity  int
}
