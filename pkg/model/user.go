package model

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}
