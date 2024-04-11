package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"test.com/be/internal/database"
)

func (apiCfg *ApiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type paramiters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	decoder := json.NewDecoder(r.Body)

	params := paramiters{}

	err := decoder.Decode(&params)
	if err != nil {
		RespondWithError(w, 400, "decoding json faled")
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
		RespondWithError(w, 400, "creating  follow feed faled")
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFollow))

}

func (apiCfg *ApiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		RespondWithError(w, 400, "getting feed faled")
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFollows(feedFollows))

}

func (apiCfg *ApiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	FeedFollowIDStr := chi.URLParam(r, "feedFollowID")

	feedFollowID, err := uuid.Parse(FeedFollowIDStr)

	if err != nil {
		RespondWithError(w, 400, "parsing id faled")
		return
	}

	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		RespondWithError(w, 400, "deleting faled")
		return
	}

	respondWithJSON(w, 200, struct{}{})
}
