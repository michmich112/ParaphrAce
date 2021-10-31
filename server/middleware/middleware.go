package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"server/context"
	"server/core/models"

	"github.com/google/uuid"
)

type createParaphraseReq struct {
	SessionToken string `json:"session_token"`
	OriginalText string `json:"text"`
}

type createUserRes struct {
	SessionToken string `json:"session_token"`
}

func CreateUser(appCtx context.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		token := uuid.New().String()
		user, err := appCtx.UserRepository.Create(models.User{SessionToken: token})
		if err != nil {
			log.Println("Could not create new user")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Could not create new user"))
			return
		}
		log.Printf("Created new user with session token: %s", user.SessionToken)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

// CreateParaphrase creates a new paraphrase in the DB
func CreateParaphrase(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var reqBody createParaphraseReq

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&reqBody)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call create paraphraseReq

	// format a response object
	res := "" //Todo implement this
	// send the response
	json.NewEncoder(w).Encode(res)
}
