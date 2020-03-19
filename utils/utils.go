package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetErrorResponse(err error, w http.ResponseWriter) {
	log.Fatalln(err.Error())

	response := struct {
		StatusCode   int    `json:"status"`
		ErrorMessage string `json:"message"`
	}{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(message)

}
func GetResponse(w http.ResponseWriter, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(response)
}
