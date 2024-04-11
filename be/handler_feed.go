package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"test.com/be/internal/database"
)

func (apiCfg *ApiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {

	type paramiters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)

	params := paramiters{}

	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, "decoding json faled")
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
		RespondWithError(w, 400, "creating feed faled")
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))

}

func (apiCfg *ApiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		RespondWithError(w, 400, "geting feeds faled")
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))

}
