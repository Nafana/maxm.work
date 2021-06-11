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
}

type App struct {
	DB        *mongo.Client
	AppConfig *AppConfig
}

func (app *AppConfig) Load() *AppConfig {
	app.Port = os.Getenv("PORT")
	app.DatabaseUsername = os.Getenv("DB_USER")
	app.DatabasePassword = os.Getenv("DB_PASSWORD")
	return app
}

func BuildDatabaseURI(user string, password string) string {
	return fmt.Sprintf("mongodb+srv://%s:%s@shortner-cluster-0.ouu65.mongodb.net/maxm-work?retryWrites=true&w=majority", user, password)
}
