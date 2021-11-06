package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/context"
	"server/core/models"
	"time"
)

type createParaphraseReq struct {
	SessionToken string `json:"session_token"`
	OriginalText string `json:"original_text"`
}

type createParaphraseRes struct {
	Result string `json:"result"`
}

// CreateParaphrase creates a new paraphrase in the DB
func CreateParaphrase(appCtx context.AppContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Context-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// START Validate request

		var reqBody createParaphraseReq
		var timestamp time.Time = time.Now()

		// decode the json request to user
		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			log.Printf("Unable to decode the request body.  %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := appCtx.UserRepository.GetBySessionToken(reqBody.SessionToken)

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if reqBody.OriginalText == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// END Validate request

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
		p, _ = appCtx.ParaphraseRespository.Update(p)

		// Call ML api
		pr, err := appCtx.ParaphraseApi.RequestParaphrase(reqBody.OriginalText)

		if err != nil {
			log.Printf("[Paraphrase Create][Error] - Error from Paraphrase Api: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// store returned text to storage
		resUri, err := appCtx.Storage.Save(fmt.Sprintf("%d-result", p.Id), pr.Paraphrase)
		if err != nil {
			log.Printf("[Paraphrase Create][Error] - Error saving result to Storage: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// update paraphrase in DB
		p.ResultFileUri = sql.NullString{
			String: resUri,
			Valid:  true,
		}
		p.StartTime = sql.NullTime{
			Time:  pr.StartTime,
			Valid: true,
		}
		p.EndTime = sql.NullTime{
			Time:  pr.EndTime,
			Valid: true,
		}

		// update in DB
		p, _ = appCtx.ParaphraseRespository.Update(p)

		// format a response object
		res := createParaphraseRes{
			Result: pr.Paraphrase,
		}

		log.Println("[Paraphrase Create][Success]")
		// send the response
		json.NewEncoder(w).Encode(res)
	}
}
