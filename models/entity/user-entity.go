package entity

import "time"

type UserEntity struct {
	Id                 uint                     `gorm:"primaryKey" json:"id"`
	Username           string                   `gorm:"not null" json:"username"`
	Password           string                   `gorm:"not null" json:"password"`
	Email              string                   `gorm:"not null" json:"email"`
	CreatedAt          time.Time                `gorm:"autoCreateTime;not null" json:"createdAt"`
	UpdatedAt          time.Time                `gorm:"autoUpdateTime;not null" json:"updatedAt"`
	UserLoginHistories []UserLoginHistoryEntity `gorm:"foreignKey:UserId"`
}

func (UserEntity) TableName() string {
	return "user_management.user"
}
