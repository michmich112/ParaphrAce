package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	aCtx "server/context"
	"server/router"

	"github.com/joho/godotenv"
)

// Middleware to inject app context
// func AppContextMiddleware(ac aCtx.AppContext) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			ctx := context.WithValue(r.Context(), "app", ac)
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		})
// 	}
// }

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("Unable to read environment variables from .env file. Verify that the environment contains all the necessary variables")
	}

	ctx := aCtx.InitAppContext()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := router.Router(ctx)
	// r.Use(AppContextMiddleware(ctx))

	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
