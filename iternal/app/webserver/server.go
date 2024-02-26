package webserver

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"web/iternal/app"
	"web/models"
)

type server struct {
	router *mux.Router
	app    app.App
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	s.router.Handle("/users", s.getUsers())
	s.router.HandleFunc("/user_create", s.userCreate()).Methods("POST")
}

func newServer(app app.App) *server {
	s := &server{
		router: mux.NewRouter(),
		app:    app,
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

func (s *server) userCreate() http.HandlerFunc {
	type request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Произошла ошибка при создании юзера")
			return
		}
		u := &models.User{
			Login:    req.Login,
			Password: req.Password,
		}
		if err := s.app.User().Create(u); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			slog.Info("error: ", err)
			fmt.Fprintf(w, "Error for create user in DB")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "пользователь успешно зарегистрирован")
	}

}
