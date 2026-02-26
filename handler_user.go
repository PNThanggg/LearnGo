package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/PNThanggg/LearnGo/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handlerGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiConfig.DB.GetUsers(r.Context())
	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("Error getting users: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseUsersToUsers(users))
}

func (apiConfig *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		APIKey string `json:"api_key"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), params.APIKey)
	if err != nil {
		responseWithError(w, 404, fmt.Sprintf("Error getting user by api_key: %v", err))
		return
	}
	responseWithJSON(w, 200, databaseUserToUser(user))
}
