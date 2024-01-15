package main

import (
	"fmt"
	"net/http"

	"github.com/nicholas-karimi/gorssaggregator/internal/auth"
	"github.com/nicholas-karimi/gorssaggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *APIConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth with error %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprint("Error getting user:", err))
			return
		}

		handler(w, r, user)
	}
}
