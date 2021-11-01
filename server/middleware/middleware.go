package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/context"
	"server/core/models"
	"time"

	"github.com/google/uuid"
)

type createParaphraseReq struct {
	SessionToken string `json:"session_token"`
	OriginalText string `json:"text"`
}

// CreateUser creates a new user in the DB
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
func CreateParaphrase(appCtx context.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		var reqBody createParaphraseReq
		var timestamp time.Time = time.Now()

		// decode the json request to user
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			log.Fatalf("Unable to decode the request body.  %v", err)
		}

		user, err := appCtx.UserRepository.GetBySessionToken(reqBody.SessionToken)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		p := models.Paraphrase{
			UserId:    user.Id,
			Timestamp: timestamp,
		}

		// call create paraphraseReq
		p, err = appCtx.ParaphraseRespository.Create(p)

		if err != nil {
			log.Printf("[Paraphrase Create][Error] - %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// store original text to storage
		uri, err := appCtx.Storage.Save(fmt.Sprintf("%d-original", p.Id), reqBody.OriginalText)

		// Update metadata with Uri of the original text
		p.OriginalFileUri = uri
		appCtx.ParaphraseRespository.Update(p)

		// Call ML api
		// store returned test to storage
		// update paraphrase in DB
		// return storage uri

		// format a response object
		//res := "" //Todo implement this
		// send the response
		//json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusOK)
	}
}
