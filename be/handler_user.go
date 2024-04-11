package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"test.com/be/internal/database"
)

func (apiCfg *ApiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type paramiters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := paramiters{}

	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, "decoding json faled")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("creating user faled : %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))

}

func (apiCfg *ApiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	// apiKey, err := auth.GetAPIKey(r.Header)
	// if err != nil {
	// 	RespondWithError(w, 402, "faled auth")
	// 	return

	// }
	// user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	// if err != nil {
	// 	RespondWithError(w, 400, "faled to get user")
	// 	return
	// }
	respondWithJSON(w, 200, databaseUserToUser(user))

}

func (apiCfg *ApiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {

	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("getting posts faled : %v", err))
		return
	}

	respondWithJSON(w, 200, databasePostsToPosts(posts))
}
