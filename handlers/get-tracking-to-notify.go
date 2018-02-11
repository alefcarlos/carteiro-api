package handlers

import (
	"net/http"

	"github.com/alefcarlos/carteiro-api/utils"

	"github.com/alefcarlos/carteiro-api/repo"
)

//GetAvailableTrackings retorna a lista de rastreios que têm novas informações
func GetAvailableTrackings(w http.ResponseWriter, r *http.Request) {
	utils.SendJSONSuccess(w, repo.GetTrackingsToNotify())
}
