package main

import (
	"fmt"
	"log"
	"net/http"
	aCtx "server/context"
	"server/router"
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

	ctx := aCtx.InitAppContext()

	port := 8080
	r := router.Router(ctx)
	// r.Use(AppContextMiddleware(ctx))

	fmt.Printf("Starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
