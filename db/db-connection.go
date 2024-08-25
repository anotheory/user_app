package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DbConnection *gorm.DB

func InitDbConnection() *gorm.DB {
	dsn := "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "test_schema.",
		},
	})
	if err != nil {
		panic(err)
	}
	DbConnection = db
	return DbConnection
}

func init() {
	InitDbConnection()
}
