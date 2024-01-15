package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/nicholas-karimi/gorssaggregator/internal/database"
)

func (apiCfg *APIConfig) handlerCreateFeedFolow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing JSON:", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Could not create feed follow:", err))
		return
	}

	// respondWithJSON(w, 200, user)
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))

}


func (apiCfg *APIConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Could not get feed follows:", err))
		return
	}

	// respondWithJSON(w, 200, user)
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollows))

}
