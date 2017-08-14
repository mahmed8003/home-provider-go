package server

import (
	"home-provider/app"
	"home-provider/controllers"
	"net/http"

	"github.com/go-chi/chi"
)

func addRoutes(ctx app.Context, router chi.Router) {

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pongno"))
	})

	router.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(userRouter chi.Router) {
			user := controllers.NewUserController(ctx)
			userRouter.Post("/", user.CreateUser)
		})
	})
}
