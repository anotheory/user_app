package healthcheck

type HealthcheckEntity struct {
	Id      uint   `gorm:"primaryKey" json:"id"`
	Message string `json:"message"`
}

func (HealthcheckEntity) TableName() string {
	return "healthcheck"
}

type HealthcheckResponse struct {
	Version string `json:"version"`
	Message string `json:"message"`
}
