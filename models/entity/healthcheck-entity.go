package entity

type HealthcheckEntity struct {
	Id      uint   `gorm:"primaryKey" json:"id"`
	Message string `json:"message"`
}

func (HealthcheckEntity) TableName() string {
	return "user_management.healthcheck"
}
