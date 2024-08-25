package repository

import (
	"user_app/db"

	"gorm.io/gorm"
)

type BaseRepository struct {
}

func (BaseRepository) Get(query *gorm.DB, receiverModels interface{}) *gorm.DB {
	results := db.DbConnection.Where(query).Find(receiverModels)
	return results
}

func (BaseRepository) GetById(id uint, receiverModel interface{}) *gorm.DB {
	results := db.DbConnection.First(receiverModel, "id = ?", id)
	return results
}

func (BaseRepository) CreateAll(models interface{}) *gorm.DB {
	results := db.DbConnection.Create(models)
	return results
}

func (BaseRepository) UpdateAll(models interface{}) *gorm.DB {
	results := db.DbConnection.Save(models)
	return results
}

func (BaseRepository) DeleteAll(models interface{}) *gorm.DB {
	results := db.DbConnection.Delete(models)
	return results
}
