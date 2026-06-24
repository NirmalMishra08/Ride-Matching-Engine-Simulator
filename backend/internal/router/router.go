package router

import "github.com/go-chi/chi"


func SetupRouter(driverHandler *handler.DriverHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Post("/drivers", driverHandler.CreateDriver)
}