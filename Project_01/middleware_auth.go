package Project_01

import (
	"fmt"
	"net/http"

	"github.com/PNThanggg/LearnGo/internal/auth"
	"github.com/PNThanggg/LearnGo/internal/database"
)

type AuthedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *apiConfig) middlewareAuth(handler AuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 404, fmt.Sprintf("Error getting user by api_key: %v", err))
			return
		}

		handler(w, r, user)
	}
}
