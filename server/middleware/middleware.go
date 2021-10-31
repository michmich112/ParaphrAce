package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type createParaphraseReq struct {
	SessionToken string `json:"session_token"`
	OriginalText string `json:"text"`
}

// CreateUser create a user in the postgres db
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
