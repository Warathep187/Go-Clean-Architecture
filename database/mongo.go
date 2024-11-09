package database

import (
	"fmt"
	"go-clean-arch/config"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDBConnection(configs *config.Config) {
	mongoUri := fmt.Sprintf("mongodb://%s:%d", configs.Db.Host, configs.Db.Port)
	err := mgm.SetDefaultConfig(&mgm.Config{
		CtxTimeout: time.Second * 3,
	},
		configs.Db.DBName,
		options.Client().ApplyURI(mongoUri),
		options.Client().SetAuth(options.Credential{
			Username: configs.Db.User,
			Password: configs.Db.Password,
		}),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Set up MongoDB connection")
}
