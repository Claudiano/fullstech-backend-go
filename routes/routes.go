package routes

import (
	"fullstech-backend-go/controllers"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func InitServer() {

	routes := chi.NewRouter()

	trabalhoController := controllers.TrabalhoController{}
	lembreteController := controllers.LembreteController{}

	routes.Route("/trabalho", func(r chi.Router) {
		r.Get("/", trabalhoController.GetAllTrabalho)
		r.Post("/", trabalhoController.PostTrabalho)
		r.Delete("/{id}", trabalhoController.DeleteTrabalho)
	})

	routes.Route("/lembrete", func(r chi.Router) {
		r.Get("/", lembreteController.GetAllLembrete)
		r.Post("/", lembreteController.PostLembrete)
		r.Delete("/{id}", lembreteController.DeleteLembrete)
	})

	log.Println("Ip servidor: http://localhost:8080")

	http.ListenAndServe(":8080", routes)

}
