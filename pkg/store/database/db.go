package database

import (
	user_models "github.com/JacobRWebb/InventoryManagement/pkg/models/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func open(dbName string) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(dbName), &gorm.Config{})
}

func MustOpen(dbName string) *gorm.DB {
	if dbName == "" {
		dbName = "default.db"
	}

	db, err := open(dbName)

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&user_models.User{})

	if err != nil {
		panic(err)
	}

	return db
}
