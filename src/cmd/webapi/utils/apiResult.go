package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//APIResultError modelo padrão para respostas de erro da API
type APIResultError struct {
	Message string `json:"message"`
}

//APIResultOk modelo padrão para reposta de sucesso da API
type APIResultOk struct {
	Result interface{} `json:"result"`
}

//SendJSONSuccess auxiliar para respostas de 200 sucesso
func SendJSONSuccess(w http.ResponseWriter, object interface{}) {
	result := APIResultOk{
		Result: object,
	}

	sendJSONReponseOK(w, result)
}

//SendJSONBadRequest auxiliar para respotas 400
func SendJSONBadRequest(w http.ResponseWriter, message string) {
	sendJSONReponseNotOK(w, message, http.StatusBadRequest)
}

//SendJSONConfict auxiliar para respotas 409
func SendJSONConfict(w http.ResponseWriter, message string) {
	sendJSONReponseNotOK(w, message, http.StatusConflict)
}

//SendJSONNotFound auxiliar para respotas 404
func SendJSONNotFound(w http.ResponseWriter, message string) {
	sendJSONReponseNotOK(w, message, http.StatusNotFound)
}

func sendJSONReponseOK(w http.ResponseWriter, object interface{}) {
	js, err := json.Marshal(object)
	if err != nil {
		sendJSONReponseNotOK(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func sendJSONReponseNotOK(w http.ResponseWriter, message string, statusCode int) {
	result := APIResultError{
		Message: message,
	}

	js, _ := json.Marshal(result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(js)
}

//ReadBody Lê o corpo e retorna um array de bytes
func ReadBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return nil, err
	}

	return body, nil
}
