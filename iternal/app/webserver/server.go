package webserver

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

type server struct {
	router *mux.Router
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.Handle("/users", s.getUsers())
}

func newServer() *server {
	s := &server{
		router: mux.NewRouter(),
	}
	s.configureRouter()
	return s
}

func (s *server) getUsers() http.HandlerFunc {
	slog.Info("use method getUsers")
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Ваш айпи зафиксирован")
	}
}
