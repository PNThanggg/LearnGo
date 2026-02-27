package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PNThanggg/LearnGo/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateFeeds(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	feed, err := apiConfig.DB.CreateNewFeed(r.Context(), database.CreateNewFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedToFeed(feed))
}
