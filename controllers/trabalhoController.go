package controllers

import (
	"encoding/json"
	"fmt"
	"fullstech-backend-go/models"
	"fullstech-backend-go/services"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type TrabalhoController struct{}

var trabalhoService services.TrabalhoService

func (TrabalhoController) GetAllTrabalho(w http.ResponseWriter, r *http.Request) {
	trabalhos := trabalhoService.GetAllTrabalhos()

	message, _ := json.Marshal(trabalhos)

	w.WriteHeader(http.StatusAccepted)
	w.Write(message)
}

func (TrabalhoController) PostTrabalho(w http.ResponseWriter, r *http.Request) {
	var trabalho models.Trabalho

	err := json.NewDecoder(r.Body).Decode(&trabalho)
	if err != nil {
		log.Fatal(err)
	} else {

		fmt.Fprintf(w, "%s", trabalhoService.SaveTrabalho(trabalho))
	}

}

func (TrabalhoController) DeleteTrabalho(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	fmt.Fprintf(w, "%s", trabalhoService.DeleteTrabalho(id))

}
