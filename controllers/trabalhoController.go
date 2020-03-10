package controllers

import (
	"fmt"
	"fullstech-backend-go/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type TrabalhoController struct{}

var trabalhoService services.TrabalhoService

func (TrabalhoController) GetAllTrabalho(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", trabalhoService.GetAllTrabalhos())
}

func (TrabalhoController) PostTrabalho(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", trabalhoService.SaveTrabalho())

}

func (TrabalhoController) DeleteTrabalho(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	fmt.Fprintf(w, "%s", trabalhoService.DeleteTrabalho(id))

}
