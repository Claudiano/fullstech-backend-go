package routes

import (
	"fullstech-backend-go/controllers"
	"net/http"

	"github.com/go-chi/chi"
)

func InitServer() {

	routes := chi.NewRouter()

	trabalhoController := controllers.TrabalhoController{}

	routes.Route("/trabalho", func(r chi.Router) {
		r.Get("/", trabalhoController.GetAllTrabalho)
		r.Post("/", trabalhoController.PostTrabalho)
		r.Delete("/{id}", trabalhoController.DeleteTrabalho)
	})

	http.ListenAndServe(":8080", routes)

}
