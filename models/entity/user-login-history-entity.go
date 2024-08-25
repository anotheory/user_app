package entity

import "time"

type UserLoginHistoryEntity struct {
	Id      uint       `gorm:"primaryKey" json:"id"`
	UserId  uint       `gorm:"not null" json:"userId"`
	LoginAt time.Time  `gorm:"autoCreateTime;not null" json:"loginAt"`
	User    UserEntity `gorm:"foreignKey:UserId"`
}

func (UserLoginHistoryEntity) TableName() string {
	return "user_management.user_login_history"
}
