package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Trying to load environment variable file if one exists
	err := godotenv.Load()
	if err != nil {
		log.WithError(err).Warn("Environment file failed to load")
	}

	// Creating an application state
	config := (&AppConfig{}).Load()
	app := &App{AppConfig: config}

	// Connecting to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	app.DB, err = mongo.Connect(ctx, options.Client().ApplyURI(BuildDatabaseURI(
		config.DatabaseUsername, config.DatabasePassword)))
	if err != nil {
		log.WithError(err).Fatal("Failed to establish connection with database")
	}

	// Create router, and initialize routes
	r := mux.NewRouter()
	r.HandleFunc("/{slug}", app.Redirect).Methods(http.MethodGet)

	log.Info(fmt.Sprintf("Started server on port %s", app.AppConfig.Port))
	log.Fatal(http.ListenAndServe(":"+app.AppConfig.Port, r))
}
