package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const validToken = "secret"

func main() {
	server := SetupServer()
	if err := http.ListenAndServe(":8080", server); err != nil {
		fmt.Printf("error on http.ListenAndServe; error=%s", err.Error())
	}
}

func SetupServer() *chi.Mux {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(AuthMiddleware)
			r.HandleFunc("/secure", secureHandler)
		})
		r.Get("/hello", helloHandler)
	})
	return r
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func secureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are authorized!")
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Auth-Token")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if token != validToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
