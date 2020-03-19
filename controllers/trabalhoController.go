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

type TrabalhoController struct{}

var trabalhoService services.TrabalhoService

func (TrabalhoController) GetAllTrabalho(w http.ResponseWriter, r *http.Request) {
	response := trabalhoService.GetAllTrabalhos()

	utils.GetResponse(w, response)
}

func (TrabalhoController) PostTrabalho(w http.ResponseWriter, r *http.Request) {
	var trabalho models.Trabalho

	err := json.NewDecoder(r.Body).Decode(&trabalho)
	if err != nil {
		log.Fatal(err)
		utils.GetErrorResponse(err, w)
	} else {

		response := trabalhoService.SaveTrabalho(trabalho)

		utils.GetResponse(w, map[string]interface{}{"message": response})
	}

}

func (TrabalhoController) DeleteTrabalho(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	response := trabalhoService.DeleteTrabalho(id)

	utils.GetResponse(w, map[string]interface{}{"message": response})

}
