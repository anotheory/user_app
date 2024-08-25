package repository

import "gorm.io/gorm"

type IBaseRepository interface {
	Get(*gorm.DB, interface{}) *gorm.DB
	GetById(uint, interface{}) *gorm.DB
	CreateAll(interface{}) *gorm.DB
	UpdateAll(interface{}) *gorm.DB
	DeleteAll(interface{}) *gorm.DB
}

func Get(repo IBaseRepository, query *gorm.DB, receiverModels []interface{}) *gorm.DB {
	return repo.Get(query, receiverModels)
}

func GetById(repo IBaseRepository, id uint, receiverModel interface{}) *gorm.DB {
	return repo.GetById(id, receiverModel)
}

func CreateAll(repo IBaseRepository, models []interface{}) *gorm.DB {
	return repo.CreateAll(models)
}

func UpdateAll(repo IBaseRepository, models []interface{}) *gorm.DB {
	return repo.UpdateAll(models)
}

func DeleteAll(repo IBaseRepository, models []interface{}) *gorm.DB {
	return repo.DeleteAll(models)
}
