package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"server/context"
	"server/core/models"

	"github.com/google/uuid"
)

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
