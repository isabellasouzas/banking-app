package http

import (
	"net/http"

	"banking-app/app"
	"banking-app/app/gateway/http/account"
	"banking-app/app/gateway/http/rest"
	"banking-app/app/gateway/http/transfer"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Router     *chi.Mux
	accHandler account.Handler
	trHandler  transfer.Handler
}

func (s *Server) Serve(address string) error {
	return http.ListenAndServe(address, s.Router)
}

func (s *Server) AssignRoutes() {

	// account routes
	s.Router.Route("/api/v1", func(r chi.Router) {
		r.Route("/accounts", func(r chi.Router) {
			r.Group(func(r chi.Router) {
				r.Method("GET", "/", rest.Handler(s.accHandler.GetBalance))
			})
			r.Method("POST", "/", rest.Handler(s.accHandler.Create))
		})

		// transfer routes
		r.Route("/transfers", func(r chi.Router) {
			r.Method("POST", "/", rest.Handler(s.trHandler.SendMoney))
			r.Method("GET", "/", rest.Handler(s.trHandler.List))
		})
	})
}

func NewServer(application *app.Application) *Server {

	srv := &Server{
		Router:     chi.NewRouter(),
		accHandler: account.NewHandler(application.AccUsecase),
		trHandler:  transfer.NewHandler(application.TrUsecase),
	}

	srv.AssignRoutes()
	return srv
}
