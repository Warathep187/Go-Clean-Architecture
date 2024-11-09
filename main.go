package main

import (
	"go-clean-arch/config"
	"go-clean-arch/database"
	"go-clean-arch/server"
)

func main() {
	conf := config.GetConfig()
	database.InitMongoDBConnection(conf)
	server.NewGinServer(conf).Start()
}
