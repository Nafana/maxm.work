package main

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

type AppConfig struct {
	Port             string
	DatabaseUsername string
	DatabasePassword string
	DatabaseHost     string
	DatabaseInitial  string
}

type App struct {
	DB        *mongo.Client
	AppConfig *AppConfig
}

func (app *AppConfig) Load() *AppConfig {
	app.Port = os.Getenv("PORT")
	app.DatabaseUsername = os.Getenv("DB_USER")
	app.DatabasePassword = os.Getenv("DB_PASSWORD")
	app.DatabaseHost = os.Getenv("DB_HOST")
	app.DatabaseInitial = os.Getenv("DB_INITIAL_DB")
	return app
}

func BuildDatabaseURI(user string, password string, host string, db string) string {
	return fmt.Sprintf("mongodb+srv://%s:%s@%s/%s?retryWrites=true&w=majority",
		user, password, host, db)
}
