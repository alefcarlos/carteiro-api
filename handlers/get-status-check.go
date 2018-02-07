package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/utils"
)

//GetStatus retorna a lista de rastreios que têm novas informações
func GetStatus(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONSuccess(w, "tudo na paz, irmão!")
}
