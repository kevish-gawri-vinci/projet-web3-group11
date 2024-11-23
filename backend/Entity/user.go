package entity

type User struct {
	ID       int `gorm:"primaryKey"`
	Username string
	Password string
	IsAdmin  bool
}
