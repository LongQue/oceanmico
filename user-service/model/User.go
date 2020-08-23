package model

type User struct {
	Username string `gorm:"type:varchar(255);unique_index"`
	Password string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(100)"`
	Channel int64 `gorm:"type:tinyint(2)"`
}
