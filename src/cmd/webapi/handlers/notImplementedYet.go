package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/src/cmd/webapi/utils"
)

//NotImplementedYet apenas para testes
func NotImplementedYet(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONBadRequest(w, "Calma, gafanhoto, esse método ainda não foi implementado!")
}
