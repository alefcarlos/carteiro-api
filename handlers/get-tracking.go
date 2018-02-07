package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/utils"
	"github.com/julienschmidt/httprouter"

	"github.com/alefcarlos/carteiro-api/repo"
)

//GetAvailableTrackings retorna a lista de rastreios que têm novas informações
func GetAvailableTrackings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	utils.SendJSONSuccess(w, repo.GetTrackings())
}
