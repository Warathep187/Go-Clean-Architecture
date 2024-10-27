package main

import (
	"go-clean-arch/config"
	"go-clean-arch/database"
	"go-clean-arch/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewDatabase(conf)
	server.NewServer(conf, db).Start()
}
