package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/utils"
	"github.com/julienschmidt/httprouter"
)

//NotImplementedYet apenas para testes
func NotImplementedYet(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.SendJSONBadRequest(w, "Calma, gafanhoto, esse método ainda não foi implementado!")
}
