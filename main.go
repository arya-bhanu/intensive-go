package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const validToken = "secret"

func main() {
	fmt.Println("Hello World")
	server := SetupServer()
	if err := http.ListenAndServe(":3000", server); err != nil {
		fmt.Printf("error on http.ListenAndServe; error=%s", err.Error())
	}
}

type Server struct {
	ChiMuxRouter *chi.Mux
}

func SetupServer() *Server {
	r := chi.NewRouter()
	// protected route
	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)
		r.HandleFunc("/secure", secureHandler)
	})
	r.Group(func(r chi.Router) {
		r.Get("/hello", helloHandler)
	})
	return &Server{
		ChiMuxRouter: r,
	}
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

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ChiMuxRouter.ServeHTTP(w, r)
}
