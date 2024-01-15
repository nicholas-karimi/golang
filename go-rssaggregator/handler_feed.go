package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nicholas-karimi/gorssaggregator/internal/database"
)

func (apiCfg *APIConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing JSON:", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error creating user:", err))
		return
	}

	// respondWithJSON(w, 200, user)
	respondWithJSON(w, 201, databaseFeedToFeed(feed))

}

func (apiCfg *APIConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprint("ECouldnt get feeds from the database:", err))
		return
	}

	// respondWithJSON(w, 200, user)
	respondWithJSON(w, 201, databaseFeedsToFeeds(feeds))

}
