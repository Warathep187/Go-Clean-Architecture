package database

import (
	"fmt"
	"sync"

	"go-clean-arch/config"
	"go-clean-arch/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewDatabase(conf *config.Config) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			conf.Db.Host,
			conf.Db.User,
			conf.Db.Password,
			conf.Db.DBName,
			conf.Db.Port,
			conf.Db.SSLMode,
			conf.Db.TimeZone,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		createNonExistingTable(db)

		dbInstance = &postgresDatabase{Db: db}
	})

	return dbInstance
}

func createNonExistingTable(db *gorm.DB) {
	if !db.Migrator().HasTable(&entities.Blog{}) {
		err := db.Migrator().CreateTable(&entities.Blog{})
		if err != nil {
			panic("failed to create Blog table")
		}
	}
	if !db.Migrator().HasTable(&entities.User{}) {
		err := db.Migrator().CreateTable(&entities.User{})
		if err != nil {
			panic("failed to create User table")
		}
	}
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
