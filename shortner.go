package main

import (
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ShortnerDatabaseName  = "maxm-work"
	ShortnerURLCollection = "shortner"
)

func (app *App) Redirect(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	s := &Slug{Name: params["slug"]}

	// Validate slug before attempting to find objects in the database
	if err := s.ValidateSlugLength(); err != nil {
		log.WithError(err).Warn(fmt.Sprintf("Client requested slug with improper length=%d", len(s.Name)))
		w.Write([]byte("400 - Invalud slug length"))
		return
	}
	if err := s.ValidateSlugFormat(); err != nil {
		log.WithError(err).Warn(fmt.Sprintf("Client requested slug=\"%s\" with improper format", s.Name))
		w.Write([]byte("400 - Invalid slug format"))
		return
	}

	collection := app.DB.Database(ShortnerDatabaseName).Collection(ShortnerURLCollection)
	filter := bson.M{"name": s.Name}
	err := collection.FindOne(r.Context(), filter).Decode(&s)
	if err != nil {
		// Document wasn't found in the database
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.WithError(err).Warn(fmt.Sprintf("Failed finding document for slug=\"%s\" in database", s.Name))
			w.Write([]byte("Slug did not exist"))
			return
		}

		// Handle an unknown error from the db
		log.WithError(err).Error(fmt.Sprintf("Failed querying database for slug=\"%s\" document", s.Name))
		w.Write([]byte("500 - Failed to process slug request"))
		return
	}

	http.Redirect(w, r, s.Redirect, http.StatusSeeOther)
	log.Info(fmt.Sprintf("Redirected client %s -> %s", s.Name, s.Redirect))
}
