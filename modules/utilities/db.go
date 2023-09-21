package utilities

import (
	"clawer/models/migrate"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Open() *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}
	db, err := gorm.Open(sqlite.Open("db/sqlite.db?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		dbInstance = nil
		panic("failed to connect database")
	}

	migrate.AutoMigration(db)

	dbInstance = db
	return dbInstance
}
