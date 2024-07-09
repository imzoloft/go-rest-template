package main

import (
	"log"

	"github.com/imzoloft/go-rest-api/cmd/api"
	"github.com/imzoloft/go-rest-api/config"
	"github.com/imzoloft/go-rest-api/database"
)

func main() {
	db := database.NewSQLDatabase(config.Database{
		Host:     config.Env.Database.Host,
		Port:     config.Env.Database.Port,
		User:     config.Env.Database.User,
		Password: config.Env.Database.Password,
		DbName:   config.Env.Database.DbName,
	})

	if err := database.HealthCheck(db); err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
