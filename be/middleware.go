package main

import (
	"net/http"

	"test.com/be/internal/auth"
	"test.com/be/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *ApiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			RespondWithError(w, 402, "faled auth")
			return

		}
		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			RespondWithError(w, 400, "faled to get user")
			return
		}
		handler(w, r, user)
	}
}
