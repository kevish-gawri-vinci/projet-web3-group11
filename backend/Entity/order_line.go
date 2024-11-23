package entity

type OrderLine struct {
	OrderId   int `gorm:"primaryKey"`
	ArticleId int `gorm:"primaryKey"`
	Quantity  int
}
