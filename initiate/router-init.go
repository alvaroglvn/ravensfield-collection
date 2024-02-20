package initiate

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func RouterInit() *chi.Mux {
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{}))

	return router
}
