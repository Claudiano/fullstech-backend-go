package controllers

import (
	"fmt"
	"fullstech-backend-go/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type LembreteController struct{}

var lembreteService services.LembreteService

func (LembreteController) GetAllLembrete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", lembreteService.GetAllLembretes())
}

func (LembreteController) PostLembrete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", lembreteService.SaveTrabalho())

}

func (LembreteController) DeleteLembrete(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	fmt.Fprintf(w, "%s", lembreteService.DeleteTrabalho(id))

}
