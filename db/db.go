package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"tweak_twitter/pkg/model"
)

func DB(connectionString string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Run DB migrations
	err = db.AutoMigrate(&model.Tracker{}, &model.Follower{})
	if err != nil {
		return nil, err
	}

	return db, err
}
