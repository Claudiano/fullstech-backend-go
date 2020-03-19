package controllers

import (
	"encoding/json"
	"fullstech-backend-go/models"
	"fullstech-backend-go/services"
	"fullstech-backend-go/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type LembreteController struct{}

var lembreteService services.LembreteService

func (LembreteController) GetAllLembrete(w http.ResponseWriter, r *http.Request) {
	response := lembreteService.GetAllLembretes()

	utils.GetResponse(w, response)

}

func (LembreteController) PostLembrete(w http.ResponseWriter, r *http.Request) {
	var lembrete models.Lembrete

	err := json.NewDecoder(r.Body).Decode(&lembrete)
	if err != nil {
		log.Fatal(err)
		utils.GetErrorResponse(err, w)
	} else {

		response := lembreteService.SaveLembrete(lembrete)

		utils.GetResponse(w, map[string]interface{}{"message": response})
	}

}

func (LembreteController) DeleteLembrete(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	response := lembreteService.DeleteLembrete(id)

	utils.GetResponse(w, map[string]interface{}{"message": response})

}
